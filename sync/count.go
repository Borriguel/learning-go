package sync

import "sync"

type Count struct {
	mu    sync.Mutex
	value int
}

func (c *Count) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Count) Value() int {
	return c.value
}
