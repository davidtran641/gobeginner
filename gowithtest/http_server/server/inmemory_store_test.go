package server

import (
	"testing"

	"github.com/davidtran641/gobeginner/gowithtest/http_server/utils/test"
)

func TestStore(t *testing.T) {
	store := NewInMemoryPlayerStore()

	player := "Julia"
	store.RecordScore(player)
	store.RecordScore(player)

	test.AssertEqual(t, 2, store.GetPlayerScore(player))
	test.AssertEqual(t, 0, store.GetPlayerScore("any"))
}
