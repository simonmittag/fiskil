package fiskil

import (
	"sync"
)

type Sequence struct {
	mu       sync.Mutex
	lastUsed int64
}

func (s *Sequence) Next() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.lastUsed++
	return s.lastUsed
}
