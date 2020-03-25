package v1

import "sync"

// Counter will increment a number
type Counter struct {
	mu    sync.Mutex
	value int
}

// NewCounter returns a new Counter
func NewCounter() *Counter {
	return new(Counter)
}

// Inc the count
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns the current count
func (c *Counter) Value() int {
	return c.value
}
