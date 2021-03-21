package pocker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	jsonContentType = "application/json"
)

// Player is datastruct of a player
type Player struct {
	Name string
	Wins int
}

// PlayerStore saves user's scores
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordScore(name string)
	GetLeague() League
}

// PlayerServer handle the server
type PlayerServer struct {
	store PlayerStore
	http.Handler
}

// NewPlayerServer returns a PlayerServer
func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := &PlayerServer{
		store: store,
	}
	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leaguageHandler))
	router.Handle("/players/", http.HandlerFunc(p.playerHandler))
	p.Handler = router
	return p
}

func (p *PlayerServer) leaguageHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(p.store.GetLeague())
}

func (p *PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request) {
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
