package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "http://orkut.com" || url == "http://myspace.com" {
		return false
	}
	return true
}

func TestWebsiteCheck(t *testing.T) {
	websites := []string{
		"http://orkut.com",
		"http://flogao.com",
		"http://myspace.com",
	}
	want := map[string]bool{
		"http://orkut.com":   false,
		"http://flogao.com":  true,
		"http://myspace.com": false,
	}
	got := WebsiteCheck(mockWebsiteChecker, websites)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want: '%v', got: '%v'", want, got)
	}
}
