package jsonspec

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonMarshal(t *testing.T) {
	bytes, err := json.Marshal(map[string]string{
		"message": "Hello, world!",
	})

	if assert.NoError(t, err) {
		assert.JSONEq(t, `{"message":"Hello, world!"}`, string(bytes))
	}
}

func TestJsonMarshalStruct(t *testing.T) {
	type data struct {
		Message string `json:"message"`
	}

	bytes, err := json.Marshal(data{
		Message: "Hello, world!",
	})

	if assert.NoError(t, err) {
		assert.JSONEq(t, `{"message":"Hello, world!"}`, string(bytes))
	}
}

func TestJsonMarshalOmitEmpty(t *testing.T) {
	type data struct {
		Message string `json:"message,omitempty"`
	}

	bytes, err := json.Marshal(data{
		Message: "",
	})

	if assert.NoError(t, err) {
		assert.JSONEq(t, `{}`, string(bytes))
	}
}

func TestJsonMarshalPointer(t *testing.T) {
	type data struct {
		Message *string `json:"message,omitempty"`
	}

	message := ""
	bytes, err := json.Marshal(data{
		Message: &message,
	})

	if assert.NoError(t, err) {
		assert.JSONEq(t, `{"message":""}`, string(bytes))
	}
}

func TestJsonMarshalNilPointer(t *testing.T) {
	type data struct {
		Message *string `json:"message,omitempty"`
	}

	bytes, err := json.Marshal(data{
		Message: nil,
	})

	if assert.NoError(t, err) {
		assert.JSONEq(t, `{}`, string(bytes))
	}
}
