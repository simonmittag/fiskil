package fiskil

import (
	"sync"
	"testing"
)

// Test the sequence by sending goroutines in parallel, getting sequence numbers. After completion, value must be
// number of attempts.
func TestSequence(t *testing.T) {
	var wg sync.WaitGroup
	var x Sequence

	for i := 0; i < 25; i++ {
		wg.Add(1)
		go func() {
			x.Next()
			wg.Done()
		}()
	}

	//async goroutines racing for the sequence wait here until last one is complete
	wg.Wait()
	var want int64 = 25
	if x.lastUsed != want {
		t.Errorf("race condition during sequence, should be %v but got %v", want, x.lastUsed)
	}
}
