package poker

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/davidtran641/gobeginner/gowithtest/http_server/utils/test"
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

			server := NewPlayerServer(store)
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

			server := NewPlayerServer(store)
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
		server := NewPlayerServer(&store)

		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())

		got := getLeagueFromResponse(t, response.Body)

		test.AssertEqual(t, "application/json", response.Header().Get("content-type"))
		test.AssertEqual(t, response.Code, http.StatusOK)
		test.AssertEqual(t, want, got)
	})
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
