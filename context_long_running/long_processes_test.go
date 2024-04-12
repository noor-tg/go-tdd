package context_long_running

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        testing.TB
}

// implement writer to simulate error handling .
// this is not exist in httptest Recorder class

type SpyResponseWriter struct {
	written bool
}

func (s SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

// the fetch method send error if context canceled
func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	// make channel for data
	data := make(chan string, 1)

	// start goroutine to semulate late response
	go func() {
		// to store response data string
		var result string
		// go throw response info
		for _, c := range s.response {
			select {
			// if there is ctx done channel response return
			case <-ctx.Done():
				// print message
				log.Println("spy store got canceled")
				return
			default:
				// sleep 10 ms and add char from response
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		// send result to data channel
		data <- result
	}()
	// to race between data sent or ctx cancel sent
	select {
	// if done sent return empty response and error
	case <-ctx.Done():
		return "", ctx.Err()
		// if data sent return it with empty error
	case res := <-data:
		return res, nil
	}
}
func TestServer(t *testing.T) {
	t.Run("make request to store normally", func(t *testing.T) {
		data := "Hello, World"
		store := SpyStore{data, t}
		srv := Server(&store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}
	})

	t.Run("stop background process if request canceled by user", func(t *testing.T) {
		data := "Hello, World"
		store := SpyStore{data, t}

		srv := Server(&store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Microsecond, cancel)
		request = request.WithContext(cancellingCtx)
		response := &SpyResponseWriter{}
		srv.ServeHTTP(response, request)

		if response.written {
			t.Error("a response should not have been written")
		}
	})
}
