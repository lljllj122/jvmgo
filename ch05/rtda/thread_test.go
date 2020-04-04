package rtda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreadCreation(t *testing.T) {
	thread := NewThread()
	assert.NotNil(t, thread)
	assert.NotNil(t, thread.stack)
}

// TODO: Will fix. fix the bug
func TestThreadOperation(t *testing.T) {
	thread := NewThread()

	frame1 := NewStackFrame(10, 10)
	frame2 := NewStackFrame(10, 10)

	thread.PushFrame(frame1)
	assert.Equal(t, frame1, thread.TopFrame())

	thread.PushFrame(frame2)
	assert.Equal(t, frame2, thread.TopFrame())

	frame := thread.PopFrame()
	assert.Equal(t, frame, frame2)
	assert.Equal(t, frame1, thread.TopFrame())
}
