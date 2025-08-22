package mocks

import (
	"bytes"
	"testing"
)

type SleeperSpy struct {
	Calls int
}

func (s *SleeperSpy) Sleep() {
	s.Calls++
}

func TestCount(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeperSpy := &SleeperSpy{}
	Count(buffer, sleeperSpy)
	result := buffer.String()
	expected := `3
2
1
Go!`
	if result != expected {
		t.Errorf("result '%s' expected '%s'", result, expected)
	}
	if sleeperSpy.Calls != 4 {
		t.Errorf("not enough sleeper calls, expected 4, result %d", sleeperSpy.Calls)
	}
}
