package rtda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackFrameCreation(t *testing.T) {
	frame := NewStackFrame(10, 10)
	assert.NotNil(t, frame)
	assert.NotEmpty(t, frame.localVars)
	assert.NotNil(t, frame.operandStack)
}
