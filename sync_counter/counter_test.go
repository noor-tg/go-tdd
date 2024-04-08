package sync_counter

import "testing"

func TestCounter(t *testing.T) {
	t.Run("increment counter 3 times  leave it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()
		assertCounter(t, counter, 3)
	})

}

func assertCounter(t testing.TB, counter Counter, want int) {
	t.Helper()
	got := counter.Value()

	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}
}
