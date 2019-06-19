package spec

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPingPongChannel(t *testing.T) {
	assert.Equal(t, PingPongChannel(), "pong")
}

func TestBlockChannel(t *testing.T) {
	assert.Equal(t, BlockChannel(), "done")
}

func TestDefer(t *testing.T) {
	assert.Equal(t, DeferredIncrement(), 0)
}
