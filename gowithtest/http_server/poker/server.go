package poker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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
	game  Game
	http.Handler
	template *template.Template
}

const (
	htmpTemplatePath = "game.html"
)

// NewPlayerServer returns a PlayerServer
func NewPlayerServer(store PlayerStore, game Game) (*PlayerServer, error) {
	p := &PlayerServer{
		store: store,
		game:  game,
	}
	tmpl, err := template.ParseFiles(htmpTemplatePath)
	if err != nil {
		return nil, fmt.Errorf("Can't create template: %v", err)
	}

	p.template = tmpl

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leaguageHandler))
	router.Handle("/players/", http.HandlerFunc(p.playerHandler))
	router.Handle("/game", http.HandlerFunc(p.gameHome))
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

func (p *PlayerServer) gameHome(w http.ResponseWriter, r *http.Request) {
	p.template.Execute(w, nil)
}

type playerServerWS struct {
	conn *websocket.Conn
}

var wsUpgradder = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func newPlayerServerWS(w http.ResponseWriter, r *http.Request) (*playerServerWS, error) {
	conn, err := wsUpgradder.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("Can't upgrade connection to websocket %v", err)
		return nil, err
	}
	return &playerServerWS{conn}, nil
}

func (p *playerServerWS) WaitForMsg() string {
	_, msg, err := p.conn.ReadMessage()
	if err != nil {
		log.Fatalf("Can't load message %v", err)
		return ""
	}
	return string(msg)
}

func (p *PlayerServer) webSocket(w http.ResponseWriter, r *http.Request) {

	ws, err := newPlayerServerWS(w, r)
	if err != nil {
		log.Fatalf("Cant create playerserver ws %v", err)
		return
	}

	numberOfPlayerMsg := ws.WaitForMsg()
	numberOfPlayer, err := strconv.Atoi(string(numberOfPlayerMsg))
	if err != nil {
		log.Fatalf("Can't read number of players: %v", err)
		return
	}

	// TODO: Handle alert
	p.game.Start(numberOfPlayer, ioutil.Discard)

	winner := ws.WaitForMsg()
	p.game.Finish(winner)
}
