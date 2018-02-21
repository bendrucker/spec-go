package math

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
  assert.Equal(t, add(1, 2), 3, "1 + 2 = 3")
}

func TestSubtract(t *testing.T) {
  assert.Equal(t, subtract(3, 1), 2, "3 - 1 = 2")
}

func TestMultiply(t *testing.T) {
  assert.Equal(t, multiply(2, 3), 6, "2 * 3 = 6")
}

func TestDivide(t *testing.T) {
  assert.Equal(t, divide(6, 2), 3, "6 / 2 = 3")
}