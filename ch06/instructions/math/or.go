package math

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

// Boolean OR int
type IOR struct {
	base.NoOperandInstruction
}

func (inst *IOR) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 | v2
	stack.PushInt(result)
}

// Boolean OR long
type LOR struct {
	base.NoOperandInstruction
}

func (inst *LOR) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 | v2
	stack.PushLong(result)
}
