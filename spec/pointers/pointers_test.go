package pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValue(t *testing.T) {
	i := 0
	fn := func(i int) int {
		i = 1
		return i
	}

	assert.Equal(t, i, 0)
	result := fn(i)
	assert.Equal(t, i, 0)
	assert.Equal(t, result, 1)
}

func TestPointer(t *testing.T) {
	i := 0
	fn := func(i *int) int {
		*i = 1
		return *i
	}

	assert.Equal(t, i, 0)
	result := fn(&i)
	assert.Equal(t, i, 1)
	assert.Equal(t, result, 1)
}
