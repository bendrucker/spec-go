package strconv

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrconvAtoi(t *testing.T) {
	i, _ := strconv.Atoi("1")
	assert.Equal(t, 1, i)
}

func TestStrconvAtoiError(t *testing.T) {
	_, err := strconv.Atoi("a")
	assert.Error(t, err)
}

func TestStrconvItoa(t *testing.T) {
	assert.Equal(t, strconv.Itoa(1), "1")
}
