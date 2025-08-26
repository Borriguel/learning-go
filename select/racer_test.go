package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares the speed of servers, returning the address of the fastest", func(t *testing.T) {
		slowServer := createServerWithDelay(20)
		fastServer := createServerWithDelay(0)
		defer slowServer.Close()
		defer fastServer.Close()
		slowUrl := slowServer.URL
		fastUrl := fastServer.URL
		want := fastUrl
		got, err := Racer(slowUrl, fastUrl)
		if err != nil {
			t.Fatalf("didn't expect an error, but got one %q", err)
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("returns an error if the server does not respond within 10s", func(t *testing.T) {
		server := createServerWithDelay(25 * time.Millisecond)
		defer server.Close()
		_, err := Configurable(server.URL, server.URL, 5*time.Millisecond)
		if err == nil {
			t.Errorf("expected an error, but I didn't get one")
		}
	})
}

func createServerWithDelay(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
}
