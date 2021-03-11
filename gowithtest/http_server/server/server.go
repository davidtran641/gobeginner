package server

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordScore(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	return &PlayerServer{store: store}
}

// ServeHTTP ...
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		p.processScore(w, r)
	case http.MethodGet:
		p.showScore(w, r)
	}

}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
	player := p.getPlayerName(r)
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "")
		return
	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) processScore(w http.ResponseWriter, r *http.Request) {
	player := p.getPlayerName(r)
	p.store.RecordScore(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) getPlayerName(r *http.Request) string {
	return strings.TrimPrefix(r.URL.Path, "/players/")
}
