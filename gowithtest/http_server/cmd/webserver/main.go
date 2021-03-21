package main

import (
	"log"
	"net/http"

	"github.com/davidtran641/gobeginner/gowithtest/http_server/pocker"
)

const (
	dbFileName = "game.db.json"
)

func main() {
	store, close, err := pocker.NewFileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal("Can't create store")
	}

	defer close()

	s := pocker.NewPlayerServer(store)
	handler := http.HandlerFunc(s.ServeHTTP)

	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatal("Could not listen on port 5000. ", err)
	}
}
