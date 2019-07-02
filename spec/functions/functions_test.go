package functions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnMultiple(t *testing.T) {
	a, b := ReturnMultiple()
	assert.Equal(t, a, "foo")
	assert.Equal(t, b, "bar")
}

func TestCreateIncrementerClosure(t *testing.T) {
	increment := CreateIncrementerClosure()

	assert.Equal(t, increment(), 1)
	assert.Equal(t, increment(), 2)

	assert.Equal(t, CreateIncrementerClosure()(), 1)
}

func TestRecursiveFibonacci(t *testing.T) {
	assert.Equal(t, Fibonacci(0), 0)
	assert.Equal(t, Fibonacci(1), 1)
	assert.Equal(t, Fibonacci(10), 55)
}
