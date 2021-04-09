package poker_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/davidtran641/gobeginner/gowithtest/http_server/poker"
	"github.com/davidtran641/gobeginner/gowithtest/http_server/utils/test"
)

func TestGameStart(t *testing.T) {
	blindAlerter := poker.NewStubBlindAlerter()

	game := poker.NewGame(blindAlerter, dummyPlayerStore)

	game.Start(5)

	cases := []poker.Alert{
		{0 * time.Second, 100},
		{10 * time.Minute, 200},
		{20 * time.Minute, 300},
		{30 * time.Minute, 400},
		{40 * time.Minute, 500},
		{50 * time.Minute, 600},
		{60 * time.Minute, 800},
		{70 * time.Minute, 1000},
		{80 * time.Minute, 2000},
		{90 * time.Minute, 4000},
		{100 * time.Minute, 8000},
	}

	checkScheduleCase(t, cases, blindAlerter)
}

func TestScheduleUser(t *testing.T) {
	t.Run("prompt user to enter number of players", func(t *testing.T) {
		blindAlerter := poker.NewStubBlindAlerter()
		game := poker.NewGame(blindAlerter, dummyPlayerStore)
		game.Start(7)

		cases := []poker.Alert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}

		checkScheduleCase(t, cases, blindAlerter)
	})
}

func checkScheduleCase(t *testing.T, cases []poker.Alert, blindAlerter *poker.StubBlindAlerter) {
	t.Helper()

	for i, c := range cases {
		tc := c
		t.Run(fmt.Sprintf("Run for %v, %d", c.ScheduledAt, c.Amount), func(t *testing.T) {
			test.AssertEqual(t, true, i < len(blindAlerter.Alerts))
			alert := blindAlerter.Alerts[i]

			test.AssertEqual(t, tc.Amount, alert.Amount)
			test.AssertEqual(t, tc.ScheduledAt, alert.ScheduledAt)
		})
	}
}

func TestGameFinish(t *testing.T) {
	store := &poker.StubPlayerStore{}
	game := poker.NewGame(dummyBlindAlert, store)
	winner := "Julia"

	game.Finish(winner)

	test.AssertEqual(t, winner, store.WinCalls[0])
}
