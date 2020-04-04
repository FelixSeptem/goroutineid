package goroutineid

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGoroutineID(t *testing.T) {
	var (
		gID = make(map[uint64]struct{})
		ch  = make(chan uint64)
		wg  = sync.WaitGroup{}
	)
	for i := 0; i < 3000; i++ {
		wg.Add(1)
		go func(ch chan<- uint64) {
			ch <- GetGoroutineID()
			wg.Done()
		}(ch)
	}
	go func() {
		for i := 0; i < 3000; i++ {
			i := <-ch
			gID[i] = struct{}{}
		}
	}()
	wg.Wait()
	assert.Equal(t, 3000, len(gID))
}

func BenchmarkGetGoroutineID(b *testing.B) {
	var wg = sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			GetGoroutineID()
			wg.Done()
		}()
		wg.Wait()
	}
}
