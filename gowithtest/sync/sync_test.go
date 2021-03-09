package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("Inc count 3 times", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("test safely concurrenntly", func(t *testing.T) {
		counter := NewCounter()
		want := 1000

		var wg sync.WaitGroup
		wg.Add(want)
		for i := 0; i < want; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		assertCounter(t, counter, want)
	})
}

func assertCounter(t *testing.T, counter *Counter, want int) {
	if counter.Value() != want {
		t.Errorf("want %d but got %d", want, counter.Value())
	}
}
