package main

import (
	"fmt"
	"sync"
)

// ref: https://blog.golang.org/pipelines
func main() {
	simplePipeline()
	doubleSquare()
	fanOutFanIn()
	fanOutFanInResourceLeak()

	fanOutFanInResourceImproved()
}

func simplePipeline() {
	fmt.Println("\n== Simple pipeline")

	// Set up the pipeline.
	c := gen(2, 3)
	out := sq(c)

	// Consume the output.
	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9
}

func doubleSquare() {
	fmt.Println("\n== Double square")

	// Set up the pipeline and consume the output
	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n) // 16 then 81
	}
}

func fanOutFanIn() {
	fmt.Println("\n== fan-out fan-in")
	in := gen(2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	// Consume the merged output from c1 and c2.
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}

}

func fanOutFanInResourceLeak() {
	fmt.Println("\n== fan-out fan-in resource leak")
	in := gen(2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	// Consume the first value from the output.
	out := merge(c1, c2)
	fmt.Println(<-out) // 4 or 9
	return
	// Since we didn't receive the second value from out,
	// one of the output goroutines is hung attempting to send it.

}

func fanOutFanInResourceImproved() {
	fmt.Println("\n== fan-out fan-in resource improved")
	done := make(chan struct{})
	defer close(done)

	in := gen(2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	// Consume the first value from the output.
	out := mergeImproved(done, c1, c2)
	fmt.Println(<-out) // 4 or 9
	return
	// Since we didn't receive the second value from out,
	// one of the output goroutines is hung attempting to send it.

}

// merge to merge output from multiple channel
func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
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

// merge to merge output from multiple channel
func mergeImproved(done <-chan struct{}, cs ...<-chan int) <-chan int {
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

// gen first stage - Gen number
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// gen first stage - Gen number
func genImproved(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	for _, n := range nums {
		out <- n
	}
	close(out)
	return out
}

// sq second stage - square number
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
