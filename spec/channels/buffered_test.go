package channels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockingChannel(t *testing.T) {
	c := make(chan string)
	go func() {
		assert.Equal(t, <-c, "yo")
	}()

	c <- "yo"
}

func buffered() chan string {
	c := make(chan string, 2)

	c <- "hello"
	c <- "world"

	return c
}

func TestBufferedChannel(t *testing.T) {
	c := buffered()
	assert.Equal(t, <-c, "hello")
	assert.Equal(t, <-c, "world")
}
