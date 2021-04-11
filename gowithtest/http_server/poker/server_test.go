package poker

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/davidtran641/gobeginner/gowithtest/http_server/utils/test"
	"github.com/gorilla/websocket"
)

var (
	gameDummy = NewStubGame()
)

func TestGetPlayers(t *testing.T) {
	store := &StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Julia":  10,
		},
		[]string{},
		[]Player{},
	}
	cases := []struct {
		name       string
		want       string
		wantStatus int
	}{
		{
			"Pepper",
			"20",
			http.StatusOK,
		},
		{
			"Julia",
			"10",
			http.StatusOK,
		},
		{
			"Apple",
			"",
			http.StatusNotFound,
		},
	}

	for _, c := range cases {
		tc := c
		t.Run(tc.name, func(t *testing.T) {
			request := newGetScoreRequest(tc.name)
			response := httptest.NewRecorder()

			server := mustMakePlayerServer(t, store, gameDummy)
			server.ServeHTTP(response, request)

			got := response.Body.String()
			gotStatus := response.Code

			test.AssertEqual(t, tc.want, got)
			test.AssertEqual(t, tc.wantStatus, gotStatus)
		})
	}
}

func TestStoreScore(t *testing.T) {
	cases := []struct {
		name       string
		want       string
		wantStatus int
		wantCall   []string
	}{
		{
			"Pepper",
			"",
			http.StatusAccepted,
			[]string{"Pepper"},
		},
		{
			"Julia",
			"",
			http.StatusAccepted,
			[]string{"Julia"},
		},
	}

	for _, c := range cases {
		tc := c
		t.Run(tc.name, func(t *testing.T) {
			request := newPostWinRequest(tc.name)
			response := httptest.NewRecorder()

			store := &StubPlayerStore{
				map[string]int{},
				[]string{},
				[]Player{},
			}

			server := mustMakePlayerServer(t, store, gameDummy)
			server.ServeHTTP(response, request)

			got := response.Body.String()
			gotStatus := response.Code

			test.AssertEqual(t, tc.want, got)
			test.AssertEqual(t, tc.wantStatus, gotStatus)
			test.AssertEqual(t, tc.wantCall, store.WinCalls)
		})
	}
}

func TestLeague(t *testing.T) {

	t.Run("/league", func(t *testing.T) {
		want := []Player{
			{"David", 10},
			{"Julia", 19},
			{"Bean", 3},
		}
		store := StubPlayerStore{nil, nil, want}
		server := mustMakePlayerServer(t, &store, gameDummy)

		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())

		got := getLeagueFromResponse(t, response.Body)

		test.AssertEqual(t, "application/json", response.Header().Get("content-type"))
		test.AssertEqual(t, response.Code, http.StatusOK)
		test.AssertEqual(t, want, got)
	})
}

func TestGame(t *testing.T) {
	t.Run("GET /game return 200", func(t *testing.T) {
		server := mustMakePlayerServer(t, &StubPlayerStore{}, gameDummy)

		request, _ := http.NewRequest(http.MethodGet, "/game", nil)

		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		test.AssertEqual(t, http.StatusOK, response.Code)
	})

	t.Run("websocket save winner", func(t *testing.T) {
		game := NewStubGame()
		store := NewStubPlayerStore()
		winner := "Julia"
		server := httptest.NewServer(mustMakePlayerServer(t, store, game))
		defer server.Close()

		wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"
		ws := mustDialWS(t, wsURL)
		defer ws.Close()

		mustWriteWSMessage(t, ws, "3")
		mustWriteWSMessage(t, ws, winner)

		time.Sleep(10 * time.Millisecond)

		test.AssertEqual(t, 3, game.PlayerCount)
		test.AssertEqual(t, winner, game.Winner)
	})
}

func mustWriteWSMessage(t *testing.T, conn *websocket.Conn, message string) {
	t.Helper()
	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Fatalf("Couldn't send message %v", err)
	}
}

func mustDialWS(t *testing.T, url string) *websocket.Conn {
	t.Helper()

	ws, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		t.Fatalf("Couldn't open ws connection %v", err)
	}
	return ws
}

func mustMakePlayerServer(t *testing.T, store PlayerStore, game Game) *PlayerServer {
	t.Helper()

	server, err := NewPlayerServer(store, game)
	if err != nil {

		t.Fatalf("Can't create server %v", err)
	}
	return server
}

func newGetScoreRequest(player string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return request
}

func getLeagueFromResponse(t *testing.T, body io.Reader) []Player {
	t.Helper()

	var got []Player
	err := json.NewDecoder(body).Decode(&got)
	test.AssertEqual(t, nil, err)

	return got
}

func newLeagueRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
}

func newPostWinRequest(player string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
	return request
}
