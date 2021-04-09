package pocker_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/davidtran641/gobeginner/gowithtest/http_server/pocker"
	"github.com/davidtran641/gobeginner/gowithtest/http_server/utils/test"
)

var dummyBlindAlert = pocker.NewStubBlindAlerter()
var dummyPlayerStore = pocker.NewStubPlayerStore()
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

func TestScheduling(t *testing.T) {
	in := strings.NewReader("5\nJulia wins\n")

	game := pocker.NewStubGame()
	cli := pocker.NewCLI(in, dummyStdOut, game)
	cli.PlayPocker()

	test.AssertEqual(t, 5, game.PlayerCount)
	test.AssertEqual(t, "Julia", game.Winner)
}

func TestPromptUser(t *testing.T) {
	t.Run("prompt user to enter number of players", func(t *testing.T) {
		stdOut := &bytes.Buffer{}
		in := strings.NewReader("7\n")

		game := pocker.NewStubGame()
		cli := pocker.NewCLI(in, stdOut, game)
		cli.PlayPocker()

		got := stdOut.String()
		want := "Please enter the number of players: "

		test.AssertEqual(t, want, got)

		test.AssertEqual(t, 7, game.PlayerCount)

	})
}
