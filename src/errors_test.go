package spec

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPanic(t *testing.T) {
	fn := func() {
		CreatePanic("boom")
	}

	assert.PanicsWithValue(t, "boom", fn)
}
