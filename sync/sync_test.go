package main

import (
	"sync"
	"testing"
)

func newCounter() *Counter {
	return &Counter{}
}

func TestCounter(t *testing.T) {
	t.Run("Increment in security run", func(t *testing.T) {
		expectCounter := 1000
		counter := newCounter()

		var wg sync.WaitGroup
		wg.Add(expectCounter)

		for i := 0; i < expectCounter; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}

		wg.Wait()

		verifyCounter(t, counter, expectCounter)
	})
}

func verifyCounter(t *testing.T, result *Counter, expect int) {
	t.Helper()

	if result.Value() != expect {
		t.Errorf("result '%d', expect '%d'", result.Value(), expect)
	}
}
