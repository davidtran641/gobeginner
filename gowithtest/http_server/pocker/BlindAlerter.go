package pocker

import (
	"fmt"
	"os"
	"time"
)

// BlindAlerter to schedule an alert
type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

// BlindAlerterFunc ...
type BlindAlerterFunc func(duration time.Duration, amount int)

// ScheduleAlertAt schedule an alert after a given duration
func (a BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int) {
	a(duration, amount)
}

// StdOutAlerter print amount after a given duration
func StdOutAlerter(duration time.Duration, amount int) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(os.Stdout, "Blind is now %d\n", amount)
	})
}
