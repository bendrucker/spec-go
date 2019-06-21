package errors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorNew(t *testing.T) {
	err := errors.New("oh no")
	assert.EqualError(t, err, "oh no")
}

type codedError struct {
	code    int
	message string
}

func (err *codedError) Error() string {
	return fmt.Sprintf("Error %d: %s", err.code, err.message)
}

func TestErrorCustomType(t *testing.T) {
	err := &codedError{
		code:    404,
		message: "Not found",
	}

	assert.EqualError(t, err, "Error 404: Not found")
}
