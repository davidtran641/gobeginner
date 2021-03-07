package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countDownStart = 3
)

func Countdown(w io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(w, i)
	}

	sleeper.Sleep()
	fmt.Fprint(w, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
