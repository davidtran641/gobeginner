package main

import (
"fmt"
"sync"
"time"
)

type SafeCounter struct {
	v map[string] int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v[key]++
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("abc")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("abc"))
}