package poker

import (
	"fmt"
	"io"
	"time"
)

// BlindAlerter to schedule an alert
type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int, to io.Writer)
}

// BlindAlerterFunc ...
type BlindAlerterFunc func(duration time.Duration, amount int, to io.Writer)

// ScheduleAlertAt schedule an alert after a given duration
func (a BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int, to io.Writer) {
	a(duration, amount, to)
}

// Alerter print amount after a given duration
func Alerter(duration time.Duration, amount int, to io.Writer) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(to, "Blind is now %d\n", amount)
	})
}
