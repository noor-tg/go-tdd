package mocking

import (
	"fmt"
	"io"
	"time"
)

// interface for contract between countdown and outside timer
type Sleeper interface {
	Sleep()
}

// generic sleeper
type DefaultSleeper struct{}

// implement sleeper based on time module
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const startCounter = 3
const finalWord = "Go!"

func Countdown(w io.Writer, sleeper Sleeper) {
	for i := startCounter; i > 0; i-- {
		fmt.Fprintln(w, i)
		sleeper.Sleep()
	}
	fmt.Fprint(w, finalWord)
}
