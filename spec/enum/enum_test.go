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
	action := Create
	assert.Equal(t, "create", fmt.Sprint(action))
}
