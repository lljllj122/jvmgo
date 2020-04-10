package math

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

// Subtract double
type DSUB struct {
	base.NoOperandInstruction
}

func (inst *DSUB) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 - v2
	stack.PushDouble(result)
}

// Subtract float
type FSUB struct {
	base.NoOperandInstruction
}

func (inst *FSUB) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 - v2
	stack.PushFloat(result)
}

// Subtract int
type ISUB struct {
	base.NoOperandInstruction
}

func (inst *ISUB) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 - v2
	stack.PushInt(result)
}

// Subtract long
type LSUB struct {
	base.NoOperandInstruction
}

func (inst *LSUB) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 - v2
	stack.PushLong(result)
}
