package main

import (
	"fmt"
	"sync"
)

// Using done to close channel
// Ref: https://blog.golang.org/pipelines
func main() {
	piplineRun()
}

func piplineRun() {
	done := make(chan struct{}, 2)
	defer close(done)

	in := gen(done, 2, 3)

	c1 := sq(done, in)
	c2 := sq(done, in)

	out := merge(done, c1, c2)

	fmt.Println(<-out) // 4 or 9

	// done will be closed by the deferred call.
}

// gen first stage - Gen number
func gen(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int, len(nums))
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}()
	return out
}

// sq second stage - square number
func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()
	return out
}

// merge to merge output from multiple channel
func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			select {
			case out <- n:
			case <-done:
			}
		}
		wg.Done()
	}
	wg.Add(len(cs))

	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
