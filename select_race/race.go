package select_race

import (
	"fmt"
	"net/http"
	"time"
)

// Requirement
// 1. send http request to 2 urls
// 2. when any of request return response return it
// 3. if no response return after 10 seconds throw error

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, 10*time.Second)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	// select between channels responses
	select {
	// case when ping with a url return it
	case <-ping(a):
		return a, nil
	// case when ping with b url return it
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timeout")
	}
}

func ping(url string) chan struct{} {
	// make new channel that return empty struct
	// empty struct will not allocate any memory
	// use make so there will not be runtime error if channel value not checked
	ch := make(chan struct{})

	// call closure as goroutine
	go func() {
		http.Get(url)
		// close channel after http request finish
		close(ch)
	}()

	// return the channel value
	return ch
}
