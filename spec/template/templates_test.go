package tpl

import (
	"bytes"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

func TestTemplate(t *testing.T) {
	outputTemplate := template.New("output")
	_, _ = outputTemplate.Parse("Formatting with {{.Package}} in {{.Language}}")

	buffer := new(bytes.Buffer)
	_ = outputTemplate.Execute(buffer, map[string]string{
		"Package":  "text/template",
		"Language": "go",
	})

	assert.Equal(t, "Formatting with text/template in go", string(buffer.Bytes()))
}
