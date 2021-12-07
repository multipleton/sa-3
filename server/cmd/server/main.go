package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/multipleton/sa-3/database"
)

var httpPortNumber = flag.Int("p", 8080, "HTTP port number")

var databaseConfiguration = database.DatabaseConfiguration{
	Host:     "localhost",
	Port:     "5432",
	User:     "sa-3-user",
	Password: "sa-3-password",
	Dbname:   "sa-3",
	Sslmode:  "disable",
}

func main() {
	flag.Parse()
	server, err := InitializeApplication(HttpPortNumber(*httpPortNumber), databaseConfiguration)
	if err != nil {
		log.Println("cannot initialize application")
		log.Fatalln(err)
	}
	go server.Start()
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt)
	<-sigChannel
	server.Stop()
}
