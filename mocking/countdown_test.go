package mocking

import (
	"bytes"
	"fmt"
	"testing"
)

// stub struct to test number of calls for sleep function
type SpySleeper struct {
	Calls int
}

// implementation for sleep function
func (s *SpySleeper) Sleep() {
	s.Calls++
	fmt.Print("calls")
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	// test number of calls for sleeper
	if spySleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper : %d", spySleeper.Calls)
	}
}
