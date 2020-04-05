package rtda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJvmStackCreation(t *testing.T) {
	stack := newJvmStack(10)
	assert.NotNil(t, stack)
	assert.Equal(t, stack.maxSize, uint(10))
	assert.Equal(t, stack.size, uint(0))
}

func TestJvmStackOperation(t *testing.T) {
	stack := newJvmStack(10)

	frame1 := NewStackFrame(nil,5, 5)
	frame2 := NewStackFrame(nil,5, 5)

	stack.push(frame1)
	assert.Equal(t, frame1, stack.top())
	assert.Equal(t, stack.size, uint(1))
	stack.push(frame2)
	assert.Equal(t, frame2, stack.top())
	assert.Equal(t, stack.size, uint(2))

	frame := stack.pop()
	assert.Equal(t, frame, frame2)
	assert.Equal(t, stack.top(), frame1)
	assert.Equal(t, stack.size, uint(1))
}

func TestJvmStackPanic(t *testing.T) {
	stack := newJvmStack(1)
	frame1 := NewStackFrame(nil,5, 5)
	frame2 := NewStackFrame(nil,5, 5)

	stack.push(frame1)

	assert.Panics(t, func() {
		stack.push(frame2)
	})
}
