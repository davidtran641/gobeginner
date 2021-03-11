package server

import (
	"fmt"
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
			}

			server := &PlayerServer{store}
			server.ServeHTTP(response, request)

			got := response.Body.String()
			gotStatus := response.Code

			test.AssertEqual(t, tc.want, got)
			test.AssertEqual(t, tc.wantStatus, gotStatus)
			test.AssertEqual(t, tc.wantCall, store.winCalls)
		})
	}
}

func newGetScoreRequest(player string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return request
}

func newPostWinRequest(player string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
	return request
}

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordScore(name string) {
	s.winCalls = append(s.winCalls, name)
}
