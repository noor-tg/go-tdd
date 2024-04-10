package context_long_running

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	canceled bool
	t        testing.TB
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.canceled = true
}

func (s *SpyStore) assertNotCanceled() {
	s.t.Helper()
	if s.canceled {
		s.t.Error("it should not have canceled the store")
	}
}

func (s *SpyStore) assertCanceled() {
	s.t.Helper()
	if !s.canceled {
		s.t.Error("store was not told to cancel")
	}
}

func TestServer(t *testing.T) {
	t.Run("make request to store normally", func(t *testing.T) {
		data := "Hello, World"
		store := SpyStore{data, false, t}
		srv := Server(&store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}

		store.assertNotCanceled()
	})

	t.Run("stop background process if request canceled by user", func(t *testing.T) {
		data := "Hello, World"
		store := SpyStore{data, false, t}
		srv := Server(&store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancelingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancelingCtx)

		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		if !store.canceled {
		}
	})
}
