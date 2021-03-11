package store

import "sync"

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

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.mu.Lock()
	defer i.mu.Unlock()

	return i.store[name]
}

func (i *InMemoryPlayerStore) RecordScore(name string) {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.store[name]++
}
