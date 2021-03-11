package main

import (
	"log"
	"net/http"

	"github.com/davidtran641/gobeginner/gowithtest/http_server/server"
	"github.com/davidtran641/gobeginner/gowithtest/http_server/store"
)

func main() {
	s := server.NewPlayerServer(store.NewInMemoryPlayerStore())
	handler := http.HandlerFunc(s.ServeHTTP)

	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatal("Could not listen on port 5000. ", err)
	}
}
