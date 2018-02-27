package spec

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestPanic(t *testing.T) {
  fn := func () {
    CreatePanic("boom")
  }

  assert.PanicsWithValue(t, "boom", fn)
}
