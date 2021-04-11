package main

import (
	"log"
	"net/http"

	"github.com/davidtran641/gobeginner/gowithtest/http_server/poker"
)

const (
	dbFileName = "game.db.json"
)

func main() {
	store, close, err := poker.NewFileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal("Can't create store")
	}

	defer close()

	game := poker.NewGame(poker.BlindAlerterFunc(poker.Alerter), store)
	s, err := poker.NewPlayerServer(store, game)
	if err != nil {
		log.Fatal("Couldn't start sever.", err)
	}
	handler := http.HandlerFunc(s.ServeHTTP)

	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatal("Could not listen on port 5000. ", err)
	}
}
