package math

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

// Boolean XOR int
type IXOR struct {
	base.NoOperandInstruction
}

func (inst *IXOR) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	result := v1 ^ v2
	stack.PushInt(result)
}

// Boolean XOR long
type LXOR struct {
	base.NoOperandInstruction
}

func (inst *LXOR) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	result := v1 ^ v2
	stack.PushLong(result)
}
