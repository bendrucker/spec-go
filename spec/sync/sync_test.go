package sync

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnceIncrement(t *testing.T) {
	assert.Equal(t, OnceIncrement(0), 1)
}
