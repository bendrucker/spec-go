package spec

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestChannel(t *testing.T) {
  assert.Equal(t, PingPongChannel(), "pong") 
}