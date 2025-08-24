package mocks

import (
	"fmt"
	"io"
	"time"
)

type SleeperConfigurable struct {
	duration time.Duration
	pause    func(time.Duration)
}

func (s *SleeperConfigurable) Sleep() {
	s.pause(s.duration)
}

type Sleeper interface {
	Sleep()
}

const lastWord = "Go!"
const startCount = 3

func Count(exit io.Writer, sleeper Sleeper) {
	for i := startCount; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(exit, i)
	}
	sleeper.Sleep()
	fmt.Fprint(exit, lastWord)
}
