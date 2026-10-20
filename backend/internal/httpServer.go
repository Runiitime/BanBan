package internal

import (
	"backend/api/handlers"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type HTTPServer struct {
	bHandlers *handlers.BoardHandler
}

func NewHTTPServer(bHandlers *handlers.BoardHandler) *HTTPServer {
	return &HTTPServer{bHandlers: bHandlers}
}

func (s *HTTPServer) Start() error {
	router := mux.NewRouter()

	router.Path("/boards").Methods("POST").HandlerFunc(s.bHandlers.HandleCreateBoard)
	router.Path("/boards").Methods("GET").HandlerFunc(s.bHandlers.HandleGetAllBoards)
	router.Path("/boards/{id}").Methods("GET").HandlerFunc(s.bHandlers.HandleGetBoard)
	router.Path("/boards/{id}").Methods("DELETE").HandlerFunc(s.bHandlers.HandleDeleteBoard)

	if err := http.ListenAndServe(":8080", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			log.Error("http server closed")
			return nil
		}

		log.Error("Failed to start HTTP server:", err)
		return err
	}

	return nil
}
