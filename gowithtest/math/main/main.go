package main

import (
	"os"
	"time"

	"github.com/davidtran641/gobeginner/gowithtest/math/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
