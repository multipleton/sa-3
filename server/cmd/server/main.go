package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/multipleton/sa-3/database"
)

var httpPortNumber = flag.Int("p", 8080, "HTTP port number")

func main() {
	flag.Parse()
	server, err := InitializeApplication(HttpPortNumber(*httpPortNumber), database.DatabaseConfiguration{}) // TODO: put db config
	if err != nil {
		log.Println("cannot initialize application")
		log.Fatalln(err)
	}
	server.Start()
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt)
	<-sigChannel
	server.Stop()
}
