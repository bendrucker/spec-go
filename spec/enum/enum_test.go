package enum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Action int

const (
	Create Action = iota
	Read
	Update
	Destroy
)

var names = []string{"create", "read", "update", "destroy"}

func (a Action) String() string {
	return names[a]
}

func TestEnum(t *testing.T) {
	assert.Equal(t, "create", fmt.Sprint(Create))
	assert.Equal(t, "read", fmt.Sprint(Read))
	assert.Equal(t, "update", fmt.Sprint(Update))
	assert.Equal(t, "destroy", fmt.Sprint(Destroy))
}
