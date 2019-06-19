package jsonspec

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalMapInterface(t *testing.T) {
	encoded := []byte(`
		{"message": "Hello, world!", "i": 1}
	`)

	var decoded map[string]interface{}
	json.Unmarshal(encoded, &decoded)

	assert.Equal(t, "Hello, world!", decoded["message"].(string))
	assert.Equal(t, 1, int(decoded["i"].(float64)))
}

func TestUnmarshallStruct(t *testing.T) {
	type event struct {
		Message string `json:"message"`
		Index   int    `json:"i"`
	}
	encoded := []byte(`
		{"message": "Hello, world!", "i": 1}
	`)

	decoded := event{}
	json.Unmarshal(encoded, &decoded)

	assert.Equal(t, "Hello, world!", decoded.Message)
	assert.Equal(t, 1, decoded.Index)
}
