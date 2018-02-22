package spec

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestPingPongChannel(t *testing.T) {
  assert.Equal(t, PingPongChannel(), "pong") 
}

func TestBlockChannel(t *testing.T) {
	assert.Equal(t, BlockChannel(), "done") 
}