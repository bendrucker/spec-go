package strings

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringsContains(t *testing.T) {
	assert.True(t, strings.Contains("food", "foo"))
}

func TestStringsFields(t *testing.T) {
	assert.Equal(t, strings.Fields("a b c"), []string{"a", "b", "c"})
}

func TestStringsJoin(t *testing.T) {
	assert.Equal(t, strings.Join([]string{"a", "b", "c"}, "."), "a.b.c")
}

func TestStringsNewReader(t *testing.T) {
	assert.Implements(t, (*io.Reader)(nil), strings.NewReader("foo"))
}
