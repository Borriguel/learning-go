package _select

import (
	"fmt"
	"net/http"
	"time"
)

var limit = 10 * time.Second

func Racer(x, y string) (winner string, err error) {
	return Configurable(x, y, limit)
}

func Configurable(x, y string, limit time.Duration) (winner string, err error) {
	select {
	case <-ping(x):
		return x, nil
	case <-ping(y):
		return y, nil
	case <-time.After(limit):
		return "", fmt.Errorf("timeout exceeded for %q and %q", x, y)
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}
