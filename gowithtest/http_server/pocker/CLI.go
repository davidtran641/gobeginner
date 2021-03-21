package pocker

import (
	"bufio"
	"io"
	"strings"
)

// CLI an pocker CLI
type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
}

// NewCLI return a CLI
func NewCLI(playerStore PlayerStore, in io.Reader) *CLI {
	return &CLI{playerStore, bufio.NewScanner(in)}
}

// PlayPocker starts playing
func (cli *CLI) PlayPocker() {
	cli.in.Scan()
	winner := extractWinner(cli.in.Text())
	cli.playerStore.RecordScore(winner)
}

func extractWinner(input string) string {
	return strings.Replace(input, " wins", "", 1)
}
