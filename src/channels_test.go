package spec

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBufferedChannel(t *testing.T) {
	channel := BufferedChannel()

	assert.Equal(t, <-channel, "hello")
	assert.Equal(t, <-channel, "world")
}
