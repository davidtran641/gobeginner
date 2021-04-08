package pocker_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/davidtran641/gobeginner/gowithtest/http_server/pocker"
	"github.com/davidtran641/gobeginner/gowithtest/http_server/utils/test"
)

func TestCLI(t *testing.T) {
	playerStore := pocker.NewStubPlayerStore()
	in := strings.NewReader("Julia wins\n")
	cli := pocker.NewCLI(playerStore, in, &StubBlindAlerter{})

	cli.PlayPocker()

	want := []string{"Julia"}
	test.AssertEqual(t, want, playerStore.WinCalls)

}

func TestScheduling(t *testing.T) {
	in := strings.NewReader("Julia  wins\n")
	playerStore := pocker.NewStubPlayerStore()
	blindAlerter := &StubBlindAlerter{}

	cli := pocker.NewCLI(playerStore, in, blindAlerter)
	cli.PlayPocker()

	cases := []Alert{
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

	for i, c := range cases {
		tc := c
		t.Run(fmt.Sprintf("Run for %v, %d", c.duration, c.amount), func(t *testing.T) {
			test.AssertEqual(t, true, i < len(blindAlerter.alerts))
			alert := blindAlerter.alerts[i]

			test.AssertEqual(t, tc.amount, alert.amount)
			test.AssertEqual(t, tc.duration, alert.scheduledAt)
		})
	}

}

type Alert struct {
	scheduledAt time.Duration
	amount      int
}
type StubBlindAlerter struct {
	alerts []Alert
}

func (s *StubBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, Alert{duration, amount})
}
