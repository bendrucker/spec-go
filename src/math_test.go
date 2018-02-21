package math

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
  assert.Equal(t, add(1, 2), 3, "1 + 2 = 3")
}