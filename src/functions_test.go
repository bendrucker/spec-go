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

func TestCreateIncrementerClosure(t *testing.T) {
  increment := CreateIncrementerClosure()
  
  assert.Equal(t, increment(), 1)
  assert.Equal(t, increment(), 2)

  assert.Equal(t, CreateIncrementerClosure()(), 1)
}