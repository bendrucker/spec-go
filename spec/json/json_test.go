package jsonspec

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKvJsonUnmarshall(t *testing.T) {
	encoded := []byte(`
		{"message": "Hello, world!", "i": 1}
	`)

	var decoded map[string]interface{}
	json.Unmarshal(encoded, &decoded)

	assert.Equal(t, "Hello, world!", decoded["message"].(string))
	assert.Equal(t, 1, int(decoded["i"].(float64)))
}
