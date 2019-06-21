package reflect

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReflectString(t *testing.T) {
	typeOf := reflect.TypeOf("foo")
	assert.Equal(t, "string", typeOf.Name())
	assert.Equal(t, reflect.String, typeOf.Kind())
}
