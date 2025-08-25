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
	expected := map[string]bool{
		"http://orkut.com":   false,
		"http://flogao.com":  true,
		"http://myspace.com": false,
	}
	result := WebsiteCheck(mockWebsiteChecker, websites)
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("expected: '%v', result: '%v'", expected, result)
	}
}
