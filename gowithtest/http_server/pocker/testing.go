package pocker

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
