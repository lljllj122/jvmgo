package math

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

// Multiply double
type DMUL struct {
	base.NoOperandInstruction
}

func (self *DMUL) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 * v2
	stack.PushDouble(result)
}

// Multiply float
type FMUL struct {
	base.NoOperandInstruction
}

func (self *FMUL) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 * v2
	stack.PushFloat(result)
}

// Multiply int
type IMUL struct {
	base.NoOperandInstruction
}

func (self *IMUL) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 * v2
	stack.PushInt(result)
}

// Multiply long
type LMUL struct {
	base.NoOperandInstruction
}

func (self *LMUL) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 * v2
	stack.PushLong(result)
}
