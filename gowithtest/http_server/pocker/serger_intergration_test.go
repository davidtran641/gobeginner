package pocker

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/davidtran641/gobeginner/gowithtest/http_server/utils/test"
)

func TestRecordingWins(t *testing.T) {
	database, cleanDatabase := createTempFile(t, "[]")
	defer cleanDatabase()
	store, err := NewFileSystemPlayerStore(database)

	test.AssertEqual(t, nil, err)

	server := NewPlayerServer(store)

	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))

	test.AssertEqual(t, http.StatusOK, response.Code)
	test.AssertEqual(t, "3", response.Body.String())
}

func TestRecordWin(t *testing.T) {
	database, cleanDatabase := createTempFile(t, "[]")
	defer cleanDatabase()
	store, err := NewFileSystemPlayerStore(database)
	server := NewPlayerServer(store)

	player := "Julia"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newLeagueRequest())

	want := []Player{
		{"Julia", 3},
	}

	got := getLeagueFromResponse(t, response.Body)
	test.AssertEqual(t, want, got)
	test.AssertEqual(t, nil, err)
}
