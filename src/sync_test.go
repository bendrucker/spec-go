package spec

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOnceIncrement(t *testing.T) {
	assert.Equal(t, OnceIncrement(0), 1)
}
