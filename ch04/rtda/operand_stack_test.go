package rtda

import (
	"jvmgo/ch04/test/assert"
	"testing"
)

func TestOpStack(t *testing.T) {
	opStack := newOperandStack(10)
	var intV int32 = 10
	var floatV float32 = 1.0
	var longV int64 = 11111
	var doubleV float64 = 1.000001
	var objectRef *Object = &Object{}

	opStack.PushInt(intV)
	opStack.PushFloat(floatV)
	opStack.PushLong(longV)
	opStack.PushDouble(doubleV)
	opStack.PushRef(objectRef)

	assert.Equal(t, objectRef, opStack.PopRef(), "")
	assert.Equal(t, doubleV, opStack.PopDouble(), "")
	assert.Equal(t, longV, opStack.PopLong(), "")
	assert.Equal(t, floatV, opStack.PopFloat(), "")
	assert.Equal(t, intV, opStack.PopInt(), "")
}
