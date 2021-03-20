package main

import (
	"log"
	"net/http"
	"os"

	"github.com/davidtran641/gobeginner/gowithtest/http_server/server"
)

const (
	dbFileName = "game.db.json"
)

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("Can't create database %s: %v", dbFileName, err)
	}

	store := server.NewPlayerStore(db)
	s := server.NewPlayerServer(&store)
	handler := http.HandlerFunc(s.ServeHTTP)

	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatal("Could not listen on port 5000. ", err)
	}
}
