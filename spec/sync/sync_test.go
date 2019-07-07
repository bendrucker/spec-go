package sync

import (
	"sync"
	"testing"

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
