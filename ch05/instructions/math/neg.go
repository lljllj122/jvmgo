package math

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

// Negate double
type DNEG struct{ base.NoOperandInstruction }

func (inst *DNEG) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushDouble(-val)
}

// Negate float
type FNEG struct{ base.NoOperandInstruction }

func (inst *FNEG) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushFloat(-val)
}

// Negate int
type INEG struct{ base.NoOperandInstruction }

func (inst *INEG) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(-val)
}

// Negate long
type LNEG struct{ base.NoOperandInstruction }

func (inst *LNEG) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushLong(-val)
}
