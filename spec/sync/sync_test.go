package sync

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOnceIncrement(t *testing.T) {
	onceIncrement := func(value int) int {
		var once sync.Once

		once.Do(func() {
			value++
		})

		once.Do(func() {
			value++
		})

		return value
	}

	assert.Equal(t, onceIncrement(0), 1)
}

func TestWaitGroup(t *testing.T) {
	wg := &sync.WaitGroup{}
	count := 5
	start := time.Now()

	fn := func() {
		time.Sleep(time.Duration(100) * time.Millisecond)
		wg.Done()
	}

	wg.Add(count)
	for i := 0; i < count; i++ {
		go fn()
	}

	wg.Wait()
	assert.InEpsilon(t, time.Duration(100)*time.Millisecond, time.Since(start), 0.05)
}
