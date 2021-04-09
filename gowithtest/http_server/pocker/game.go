package pocker

import "time"

// Game to manage game play
type Game interface {
	Start(numberOfPlayer int)
	Finish(winner string)
}

// TexasGame to manage the game play
type TexasGame struct {
	alerter BlindAlerter
	store   PlayerStore
}

// NewGame to create a game
func NewGame(alerter BlindAlerter, store PlayerStore) *TexasGame {
	return &TexasGame{alerter: alerter, store: store}
}

// Start to begin the game play
func (p *TexasGame) Start(numberOfPlayer int) {
	blindIncrement := time.Duration(5+numberOfPlayer) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, v := range blinds {
		p.alerter.ScheduleAlertAt(blindTime, v)
		blindTime = blindTime + blindIncrement
	}
}

// Finish to end the game with winner
func (p *TexasGame) Finish(winner string) {
	p.store.RecordScore(winner)
}
