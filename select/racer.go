package _select

import (
	"net/http"
	"time"
)

func Racer(x, y string) (winner string) {
	durationX := measureResponseTime(x)
	durationY := measureResponseTime(y)
	if durationX < durationY {
		return x
	}
	return y
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
