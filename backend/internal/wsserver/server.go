package wsserver

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type WSServer interface {
	Start() error
	Stop() error
}

type wsSrv struct {
	srv       *http.Server
	mux       *http.ServeMux
	wsUpg     *websocket.Upgrader
	wsClients map[*websocket.Conn]struct{}
	mutex     *sync.RWMutex
	broadcast chan *wsMessage
}

func NewWsServer(addr string) WSServer {
	m := http.NewServeMux()
	return &wsSrv{
		mux: m,
		srv: &http.Server{
			Addr:    addr,
			Handler: m,
		},
		wsUpg: &websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true // разрешить все Origin — только для локальной разработки!
			},
		},
		wsClients: make(map[*websocket.Conn]struct{}),
		mutex:     &sync.RWMutex{},
		broadcast: make(chan *wsMessage),
	}
}

func (ws *wsSrv) Start() error {
	ws.mux.HandleFunc("/ws", ws.wsHandler)

	go ws.broadcastToClients()
	return ws.srv.ListenAndServe()
}

func (ws *wsSrv) Stop() error {
	fmt.Println("before", ws.wsClients)
	close(ws.broadcast)
	ws.mutex.Lock()
	for conn := range ws.wsClients {
		go func() {
			if err := conn.Close(); err != nil {
				log.Println("Error closing connection:", err)
			}
		}()
		delete(ws.wsClients, conn)
	}
	ws.mutex.Unlock()

	fmt.Println("after", ws.wsClients)
	return ws.srv.Shutdown(context.Background())
}

func (ws *wsSrv) wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := ws.wsUpg.Upgrade(w, r, nil)

	if err != nil {
		log.Println("Error upgrading to websocket:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ws.mutex.Lock()
	ws.wsClients[conn] = struct{}{}
	ws.mutex.Unlock()

	go ws.readFromClient(conn)
}

func (ws *wsSrv) readFromClient(conn *websocket.Conn) {

	for {
		msg := new(wsMessage)

		if err := conn.ReadJSON(msg); err != nil {
			wsErr, ok := err.(*websocket.CloseError)
			if !ok || wsErr.Code != websocket.CloseGoingAway {
				log.Println("Error reading from WebSocket:", err)
			}
			break
		}

		host, _, err := net.SplitHostPort(conn.RemoteAddr().String())
		if err != nil {
			log.Println("Error parsing remote address:", err)
		}

		msg.IPAddress = host
		msg.Time = time.Now().Format("15:04:05")

		ws.broadcast <- msg
	}

	ws.mutex.Lock()
	delete(ws.wsClients, conn)
	ws.mutex.Unlock()
}

func (ws *wsSrv) broadcastToClients() {
	for msg := range ws.broadcast {
		ws.mutex.RLock()
		for client := range ws.wsClients {
			func() {
				if err := client.WriteJSON(msg); err != nil {
					log.Println("Error writing to client:", err)
				}
			}()

		}
		ws.mutex.RUnlock()
	}
}
