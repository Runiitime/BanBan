package main

import (
	"backend/api/handlers"
	"backend/internal"
	"backend/internal/pgs"
	"backend/internal/pgs/queries"
	"context"

	log "github.com/sirupsen/logrus"
)

const (
	addr = "127.0.0.1:9090"
)

func main() {
	ctx := context.Context(context.Background())
	conn, err := pgs.CreateConnection(ctx)

	if err != nil {
		log.Println("Error connecting to DB:", err)
	}

	bQuery := queries.NewBoardQuery("boards", conn)
	bHandler := handlers.NewBoardHandler(bQuery)

	httpSrv := internal.NewHTTPServer(bHandler)

	log.Info("Listening on 8080")
	if err := httpSrv.Start(); err != nil {
		log.Fatalln("Failed to start HTTP server:", err)
	}

	//wsSrv := wsserver.NewWsServer(addr)
	//
	//ctx := context.Context(context.Background())
	//
	//log.Println("Starting ws server...")
	//
	//_, err := pgs.CreateConnection(ctx)
	//
	//if err != nil {
	//	log.Println("Error connecting to DB:", err)
	//}
	//
	////go func() {
	//if err := wsSrv.Start(); err != nil {
	//	log.Println("Oooops, smth went wrong", err.Error())
	//}
	//
	////}()
	//
	////time.Sleep(15 * time.Second)
	////wsSrv.Stop()
}
