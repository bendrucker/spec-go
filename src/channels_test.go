package spec

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestBufferedChannel(t *testing.T) {
  channel := BufferedChannel()

  assert.Equal(t, <-channel, "hello")
  assert.Equal(t, <-channel, "world") 
}