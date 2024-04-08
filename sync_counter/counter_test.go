package sync_counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("increment counter 3 times  leave it at 3", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}
		var wg sync.WaitGroup
		wg.Add(wantedCount)
		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()
		assertCounter(t, &counter, 1000)
	})

}

func assertCounter(t testing.TB, counter *Counter, want int) {
	t.Helper()
	got := counter.Value()

	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}
}
