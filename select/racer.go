package _select

import (
	"net/http"
)

func Racer(x, y string) (winner string) {
	select {
	case <-ping(x):
		return x
	case <-ping(y):
		return y
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
