package mocking

import (
	"fmt"
	"io"
	"time"
)

const startCounter = 3
const finalWord = "Go!"

func Countdown(w io.Writer) {
	for i := startCounter; i > 0; i-- {
		fmt.Fprintln(w, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(w, finalWord)
}
