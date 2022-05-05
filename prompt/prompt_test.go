package prompt

import (
	"bytes"
	"strings"
	"testing"

	"github.com/manifoldco/promptui"
	"github.com/stretchr/testify/assert"
)

type ClosingBuffer struct {
	*bytes.Buffer
}

func (cb ClosingBuffer) Close() error {
	return nil
}

func TestPrompt(t *testing.T) {
	var input string = "Test input"
	var reader ClosingBuffer = ClosingBuffer{
		bytes.NewBufferString(input + "\n"),
	}

	var response string = Prompt(promptui.Prompt{
		Stdin: reader,
	}, nil)
	assert.True(t, strings.EqualFold(response, input))
}
