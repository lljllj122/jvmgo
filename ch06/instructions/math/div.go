package math

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

// Divide double
type DDIV struct {
	base.NoOperandInstruction
}

func (inst *DDIV) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 / v2
	stack.PushDouble(result)
}

// Divide float
type FDIV struct {
	base.NoOperandInstruction
}

func (inst *FDIV) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 / v2
	stack.PushFloat(result)
}

// Divide int
type IDIV struct {
	base.NoOperandInstruction
}

func (inst *IDIV) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 / v2
	stack.PushInt(result)
}

// Divide long
type LDIV struct {
	base.NoOperandInstruction
}

func (inst *LDIV) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 / v2
	stack.PushLong(result)
}
