package sync

import (
	"sync"
	"testing"
)

func TestCount(t *testing.T) {
	t.Run("increment the counter 3 times result in the value 3", func(t *testing.T) {
		count := Count{}
		count.Increment()
		count.Increment()
		count.Increment()
		verifyCount(t, count, 3)
	})
	t.Run("runs concurrently safely", func(t *testing.T) {
		countWant := 1000
		count := Count{}
		var wg sync.WaitGroup
		wg.Add(countWant)
		for i := 0; i < countWant; i++ {
			go func(w *sync.WaitGroup) {
				count.Increment()
				w.Done()
			}(&wg)
		}
		wg.Wait()
		verifyCount(t, count, countWant)
	})
}

func verifyCount(t *testing.T, result Count, want int) {
	t.Helper()
	if result.Value() != want {
		t.Errorf("got %q, want %q", result.Value(), want)
	}
}
