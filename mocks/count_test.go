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
		result := buffer.String()
		expected := `3
2
1
Go!`
		if result != expected {
			t.Errorf("result '%s' expected '%s'", result, expected)
		}
	})
	t.Run("pause before each print", func(t *testing.T) {
		spyPrintSleep := &SpyCountOperations{}
		Count(spyPrintSleep, spyPrintSleep)
		expected := []string{pause, write, pause, write, pause, write, pause, write}
		if !reflect.DeepEqual(expected, spyPrintSleep.Calls) {
			t.Errorf("expected '%v', result '%v'", expected, spyPrintSleep.Calls)
		}
	})
}

func TestSleeperConfigurable(t *testing.T) {
	timePause := 5 * time.Second
	timeSpy := &TimeSpy{}
	sleeper := SleepConfigurable{timePause, timeSpy.Sleep}
	sleeper.Pause()
	if timeSpy.pauseDuration != timePause {
		t.Errorf("should have paused for %v, but paused for %v", timePause, timeSpy.pauseDuration)
	}
}

type SleeperSpy struct {
	Calls int
}
type SleepConfigurable struct {
	duration time.Duration
	pause    func(time.Duration)
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

func (s *SleepConfigurable) Pause() {
	s.pause(s.duration)
}

const write = "write"
const pause = "pause"

func (s *SleeperSpy) Sleep() {
	s.Calls++
}
