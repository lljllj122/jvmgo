package conversions

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

// Convert int to byte
// Value is popped from the operand stack, truncated to a byte, then sign-extended to an int result.
type I2B struct {
	base.NoOperandInstruction
}

func (inst *I2B) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	// truncate and sign-extended to an int
	b := int32(int8(i))
	stack.PushInt(b)
}

// Convert int to char
// Value is popped from the operand stack, truncated to char, then zero-extended to an int result.
type I2C struct {
	base.NoOperandInstruction
}

func (inst *I2C) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	// truncate and zero-extended to an int result
	c := int32(uint16(i))
	stack.PushInt(c)
}

// Convert int to short
// is popped from the operand stack, truncated to a short, then sign-extended to an int result.
type I2S struct {
	base.NoOperandInstruction
}

func (inst *I2S) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	// truncated to a short
	s := int16(i)
	//sign-extended to an int
	res := int32(s)
	stack.PushInt(res)
}

// Convert int to long
type I2L struct {
	base.NoOperandInstruction
}

func (inst *I2L) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	l := int64(i)
	stack.PushLong(l)
}

// Convert int to float
type I2F struct {
	base.NoOperandInstruction
}

func (inst *I2F) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	f := float32(i)
	stack.PushFloat(f)
}

// Convert int to double
type I2D struct {
	base.NoOperandInstruction
}

func (inst *I2D) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	d := float64(i)
	stack.PushDouble(d)
}
