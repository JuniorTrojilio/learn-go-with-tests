package main

import "sync"

// Counter is a struct to counter
type Counter struct {
	mu    sync.Mutex
	value int
}

// Inc increments value to Counter.value
func (c *Counter) Inc() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

// Value return value of Counter
func (c *Counter) Value() int {
	return c.value
}

func main() {}
