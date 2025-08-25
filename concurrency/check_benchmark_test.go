package concurrency

import (
	"testing"
	"time"
)

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}
func BenchmarkWebsiteCheck(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "randomultraurls"
	}
	for i := 0; i < b.N; i++ {
		WebsiteCheck(slowStubWebsiteChecker, urls)
	}
}
