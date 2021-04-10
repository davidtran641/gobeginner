package poker

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/gorilla/websocket"
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
	template *template.Template
}

const (
	htmpTemplatePath = "game.html"
)

// NewPlayerServer returns a PlayerServer
func NewPlayerServer(store PlayerStore) (*PlayerServer, error) {
	p := &PlayerServer{
		store: store,
	}
	tmpl, err := template.ParseFiles(htmpTemplatePath)
	if err != nil {
		return nil, fmt.Errorf("Can't create template: %v", err)
	}

	p.template = tmpl

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leaguageHandler))
	router.Handle("/players/", http.HandlerFunc(p.playerHandler))
	router.Handle("/game", http.HandlerFunc(p.game))
	router.Handle("/ws", http.HandlerFunc(p.webSocket))
	p.Handler = router
	return p, nil
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

func (p *PlayerServer) game(w http.ResponseWriter, r *http.Request) {
	p.template.Execute(w, nil)
}

var wsUpgradder = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (p *PlayerServer) webSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := wsUpgradder.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("Can't upgrade connection to websocket %v", err)
		return
	}

	_, winnerMsg, err := conn.ReadMessage()
	if err != nil {
		log.Fatalf("Can't read message: %v", err)
		return
	}
	p.store.RecordScore(string(winnerMsg))
}
