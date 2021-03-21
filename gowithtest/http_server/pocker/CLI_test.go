package pocker_test

import (
	"strings"
	"testing"

	"github.com/davidtran641/gobeginner/gowithtest/http_server/pocker"
	"github.com/davidtran641/gobeginner/gowithtest/http_server/utils/test"
)

func TestCLI(t *testing.T) {
	playerStore := pocker.NewStubPlayerStore()
	in := strings.NewReader("Julia wins\n")
	cli := pocker.NewCLI(playerStore, in)

	cli.PlayPocker()

	want := []string{"Julia"}
	test.AssertEqual(t, want, playerStore.WinCalls)

}
