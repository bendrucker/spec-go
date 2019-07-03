package structs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNamedField(t *testing.T) {
	type MyStruct struct {
		Foo string
	}

	value := MyStruct{"bar"}
	assert.Equal(t, "bar", value.Foo)
}

func TestAnonymousField(t *testing.T) {
	type Embedded struct {
		Foo string
	}

	type MyStruct struct {
		Embedded
	}

	value := MyStruct{Embedded{"bar"}}
	assert.Equal(t, "bar", value.Embedded.Foo)
	assert.Equal(t, "bar", value.Foo)
}
