package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPanic(t *testing.T) {
	fn := func() {
		panic("boom")
	}

	assert.PanicsWithValue(t, "boom", fn)
}
