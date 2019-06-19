package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, Add(1, 2), 3, "1 + 2 = 3")
}

func TestAddVariadic(t *testing.T) {
	assert.Equal(t, AddVariadic(1, 2, 3, 4), 10, "1 + 2 + 3 + 4 = 10")
}

func TestSubtract(t *testing.T) {
	assert.Equal(t, Subtract(3, 1), 2, "3 - 1 = 2")
}

func TestMultiply(t *testing.T) {
	assert.Equal(t, Multiply(2, 3), 6, "2 * 3 = 6")
}

func TestDivide(t *testing.T) {
	assert.Equal(t, Divide(6, 2), 3, "6 / 2 = 3")
}
