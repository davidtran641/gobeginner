package server

import (
	"sync"
)

// InMemoryPlayerStore an implementation of PlayerScore in memory
type InMemoryPlayerStore struct {
	mu    sync.Mutex
	store map[string]int
}

// NewInMemoryPlayerStore ...
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		mu:    sync.Mutex{},
		store: map[string]int{},
	}
}

// GetPlayerScore returns player score given player name
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.mu.Lock()
	defer i.mu.Unlock()

	return i.store[name]
}

// RecordScore save score of the player
func (i *InMemoryPlayerStore) RecordScore(name string) {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.store[name]++
}

// GetLeague return top players
func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}
