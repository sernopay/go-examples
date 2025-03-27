package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (s *SafeCounter) Inc(key string) {
	s.mu.Lock() // Lock so only one goroutine at a time can access the map c.v
	s.v[key]++
	s.mu.Unlock()
}

// Value returns the current value of the counter for the given key
func (s *SafeCounter) Value(key string) int {
	s.mu.Lock() // Lock so only one goroutine at a time can access the map c.v
	defer s.mu.Unlock()
	return s.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 100; i++ {
		go c.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
