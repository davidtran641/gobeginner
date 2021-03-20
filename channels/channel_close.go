package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(chan int, 3)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		c <- 1

		time.Sleep(5 * time.Microsecond)

		c <- 2
	}()

	done := make(chan bool)
	go func() {
		defer close(done)
		wg.Wait()
	}()

	timeout := 1 * time.Microsecond
	select {
	case <-done:
		break
	case <-time.After(timeout):
		break
	}

	// close(c)
	timeoutChan := time.After(0)
forLabel:
	for i := 0; i < 3; i++ {
		select {
		case v := <-c:
			fmt.Println(v)
			break forLabel
		case <-timeoutChan:
			fmt.Println("timeout")
		}
	}

	time.Sleep(10 * time.Microsecond)

	// panic("Print stack trace")
}
