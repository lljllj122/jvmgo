package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// Add double
type DADD struct {
	base.NoOperandInstruction
}

func (inst *DADD) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	result := v1 + v2
	stack.PushDouble(result)
}

// Add float
type FADD struct {
	base.NoOperandInstruction
}

func (inst *FADD) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 + v2
	stack.PushFloat(result)
}

// Add int
type IADD struct {
	base.NoOperandInstruction
}

func (inst *IADD) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 + v2
	stack.PushInt(result)
}

// Add long
type LADD struct {
	base.NoOperandInstruction
}

func (inst *LADD) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 + v2
	stack.PushLong(result)
}
