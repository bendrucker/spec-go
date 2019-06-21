package spawn

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecEcho(t *testing.T) {
	echo := exec.Command("echo", "hello", "world")

	output, err := echo.Output()

	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, string(output), "hello world\n")
}
