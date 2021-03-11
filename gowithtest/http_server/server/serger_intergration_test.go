package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/davidtran641/gobeginner/gowithtest/http_server/store"
	"github.com/davidtran641/gobeginner/gowithtest/http_server/utils/test"
)

func TestRecordingWins(t *testing.T) {
	store := store.NewInMemoryPlayerStore()
	server := &PlayerServer{store}

	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))

	test.AssertEqual(t, http.StatusOK, response.Code)
	test.AssertEqual(t, "3", response.Body.String())
}
