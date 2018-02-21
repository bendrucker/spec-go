package spec

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestReturnMultiple(t *testing.T) {
  a, b := ReturnMultiple()
  assert.Equal(t, a, "foo")
  assert.Equal(t, b, "bar")
}