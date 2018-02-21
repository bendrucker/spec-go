package spec

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
  assert.Equal(t, Add(1, 2), 3, "1 + 2 = 3")
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