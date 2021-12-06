package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/multipleton/sa-3/utils"
)

type HttpPortNumber int

type ApiServer struct {
	port        HttpPortNumber
	router      *mux.Router
	server      *http.Server
	controllers []utils.HTTPController
}

func (as *ApiServer) Start() {
	for _, controller := range as.controllers {
		controller.HandleRoutes(as.router)
	}
	as.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", as.port),
		Handler: as.router,
	}
	log.Printf("starting server on %d", as.port)
	err := as.server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func (as *ApiServer) Stop() {
	if as.server == nil {
		log.Println("cannot stop server, server was not started")
	}
	log.Println("gracefully shuting down the server")
	err := as.server.Shutdown(context.Background())
	if err != nil {
		log.Println("gracefull shutdown failed")
		log.Fatalln(err)
	}
}
