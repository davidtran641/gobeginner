package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	// PlayerPrompt ...
	PlayerPrompt = "Please enter the number of players: "

	// BadInputMsg ...
	BadInputMsg = "Bad value received for number of players, please try again with a number"
)

// CLI an poker CLI
type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

// NewCLI return a CLI
func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{bufio.NewScanner(in), out, game}
}

// PlayPoker starts playing
func (cli *CLI) PlayPoker() {
	fmt.Fprintf(cli.out, PlayerPrompt)

	numberOfPlayer, err := strconv.Atoi(cli.readLine())

	if err != nil {
		fmt.Fprintf(cli.out, BadInputMsg)
		return
	}

	cli.game.Start(numberOfPlayer)

	userInput := cli.readLine()
	winner := extractWinner(userInput)
	cli.game.Finish(winner)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func extractWinner(input string) string {
	return strings.Replace(input, " wins", "", 1)
}
