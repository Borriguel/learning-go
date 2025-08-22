package mocks

import (
	"bytes"
	"fmt"
	"time"
)

type SleeperStandard struct {
}

func (d *SleeperStandard) Sleep() {
	time.Sleep(1 * time.Second)
}

type Sleeper interface {
	Sleep()
}

const lastWord = "Go!"
const startCount = 3

func Count(exit *bytes.Buffer, sleeper Sleeper) {
	for i := startCount; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(exit, i)
	}
	sleeper.Sleep()
	fmt.Fprint(exit, lastWord)
}
