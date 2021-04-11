package poker

import (
	"io"
	"time"
)

// StubPlayerStore return fake PlayerStore
type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	League   League
}

// NewStubPlayerStore create a StubPlayerStore
func NewStubPlayerStore() *StubPlayerStore {
	return &StubPlayerStore{}
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.Scores[name]
	return score
}

func (s *StubPlayerStore) RecordScore(name string) {
	s.WinCalls = append(s.WinCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.League
}

type Alert struct {
	ScheduledAt time.Duration
	Amount      int
}
type StubBlindAlerter struct {
	Alerts []Alert
}

func NewStubBlindAlerter() *StubBlindAlerter {
	return &StubBlindAlerter{}
}

func (s *StubBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int, to io.Writer) {
	s.Alerts = append(s.Alerts, Alert{duration, amount})
}

type StubGame struct {
	PlayerCount int
	Winner      string
	StartCalled bool

	BlindAlert []byte
}

func NewStubGame() *StubGame {
	return &StubGame{}
}

func (p *StubGame) Start(numberOfPlayer int, alertDestination io.Writer) {
	p.StartCalled = true
	p.PlayerCount = numberOfPlayer

	alertDestination.Write(p.BlindAlert)
}
func (p *StubGame) Finish(winner string) {
	p.Winner = winner
}
