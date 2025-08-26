package mocks

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCount(t *testing.T) {
	t.Run("prints 3 until Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Count(buffer, &SpyCountOperations{})
		got := buffer.String()
		want := `3
2
1
Go!`
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
	t.Run("pause before each print", func(t *testing.T) {
		spyPrintSleep := &SpyCountOperations{}
		Count(spyPrintSleep, spyPrintSleep)
		want := []string{pause, write, pause, write, pause, write, pause, write}
		if !reflect.DeepEqual(want, spyPrintSleep.Calls) {
			t.Errorf("want '%v', got '%v'", want, spyPrintSleep.Calls)
		}
	})
}

func TestSleeperConfigurable(t *testing.T) {
	timePause := 5 * time.Second
	timeSpy := &TimeSpy{}
	sleeper := SleeperConfigurable{timePause, timeSpy.Sleep}
	sleeper.Sleep()
	if timeSpy.pauseDuration != timePause {
		t.Errorf("should have paused for %v, but paused for %v", timePause, timeSpy.pauseDuration)
	}
}

type SleeperSpy struct {
	Calls int
}

type SpyCountOperations struct {
	Calls []string
}
type TimeSpy struct {
	pauseDuration time.Duration
}

func (t *TimeSpy) Sleep(duration time.Duration) {
	t.pauseDuration = duration
}

func (s *SpyCountOperations) Sleep() {
	s.Calls = append(s.Calls, pause)
}

func (s *SpyCountOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func (s *SleeperConfigurable) Sleep() {
	s.pause(s.duration)
}

const write = "write"
const pause = "pause"

func (s *SleeperSpy) Sleep() {
	s.Calls++
}
