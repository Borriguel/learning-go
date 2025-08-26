package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := createServerWithDelay(20)
	fastServer := createServerWithDelay(0)
	defer slowServer.Close()
	defer fastServer.Close()
	slowUrl := slowServer.URL
	fastUrl := fastServer.URL
	want := fastUrl
	got := Racer(slowUrl, fastUrl)
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func createServerWithDelay(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
}
