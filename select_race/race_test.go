package select_race

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	// so mocking the servers with testing code is to test behavior not implementation
	// and it will make tests faster and more reliable
	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(10 * time.Millisecond)

	// close server in end of test
	defer slowServer.Close()
	defer fastServer.Close()

	slowUrl := slowServer.URL
	fastUrl := fastServer.URL

	want := fastUrl
	got := Racer(slowUrl, fastUrl)

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	// start testing server
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
