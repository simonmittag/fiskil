package fiskil

import (
	"sync"
)

// Sequence is a wrapper type for an int, controlled by a Mutex
type Sequence struct {
	mu       sync.Mutex
	lastUsed int64
}

// Gets the next sequence number in a threadsafe way
func (s *Sequence) Next() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.lastUsed++
	return s.lastUsed
}
