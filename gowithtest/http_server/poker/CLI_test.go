package poker_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/davidtran641/gobeginner/gowithtest/http_server/poker"
	"github.com/davidtran641/gobeginner/gowithtest/http_server/utils/test"
)

var dummyBlindAlert = poker.NewStubBlindAlerter()
var dummyPlayerStore = poker.NewStubPlayerStore()
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

func TestScheduling(t *testing.T) {
	in := strings.NewReader("5\nJulia wins\n")

	game := poker.NewStubGame()
	cli := poker.NewCLI(in, dummyStdOut, game)
	cli.PlayPoker()

	test.AssertEqual(t, 5, game.PlayerCount)
	test.AssertEqual(t, "Julia", game.Winner)
}

func TestPromptUser(t *testing.T) {
	t.Run("prompt user to enter number of players", func(t *testing.T) {
		stdOut := &bytes.Buffer{}
		in := strings.NewReader("7\n")

		game := poker.NewStubGame()
		cli := poker.NewCLI(in, stdOut, game)
		cli.PlayPoker()

		got := stdOut.String()
		want := "Please enter the number of players: "

		test.AssertEqual(t, want, got)

		test.AssertEqual(t, 7, game.PlayerCount)

	})

	t.Run("Wrong number", func(t *testing.T) {
		stdOut := &bytes.Buffer{}
		in := strings.NewReader("apple\n")

		game := poker.NewStubGame()
		cli := poker.NewCLI(in, stdOut, game)
		cli.PlayPoker()

		test.AssertEqual(t, false, game.StartCalled)

		want := poker.PlayerPrompt + poker.BadInputMsg
		test.AssertEqual(t, want, stdOut.String())
	})
}
