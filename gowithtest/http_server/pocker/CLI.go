package pocker

import (
	"bufio"
	"io"
	"strings"
	"time"
)

// CLI an pocker CLI
type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
	alerter     BlindAlerter
}

// NewCLI return a CLI
func NewCLI(playerStore PlayerStore, in io.Reader, alerter BlindAlerter) *CLI {
	return &CLI{playerStore, bufio.NewScanner(in), alerter}
}

// PlayPocker starts playing
func (cli *CLI) PlayPocker() {

	cli.scheduleBlindAlerts()

	cli.in.Scan()
	winner := extractWinner(cli.in.Text())
	cli.playerStore.RecordScore(winner)
}

func (cli *CLI) scheduleBlindAlerts() {
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, v := range blinds {
		cli.alerter.ScheduleAlertAt(blindTime, v)
		blindTime = blindTime + 10*time.Minute
	}
}

func extractWinner(input string) string {
	return strings.Replace(input, " wins", "", 1)
}
