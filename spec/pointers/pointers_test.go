package pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValue(t *testing.T) {
	i := 0
	fn := func(input int) int {
		assert.Equal(t, input, 0)
		input = 1
		return input
	}

	assert.Equal(t, 0, i)
	result := fn(i)
	assert.Equal(t, 0, i)
	assert.Equal(t, 1, result)
}

func TestPointer(t *testing.T) {
	i := 0
	fn := func(input *int) int {
		assert.Equal(t, 0, *input)
		*input = 1
		return *input
	}

	assert.Equal(t, i, 0)
	result := fn(&i)
	assert.Equal(t, i, 1)
	assert.Equal(t, result, 1)
}
