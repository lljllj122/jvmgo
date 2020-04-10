package conversions

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

/*
D2F - convert stack top double value to float
*/
type D2F struct {
	base.NoOperandInstruction
}

func (inst *D2F) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	stack.PushFloat(float32(d))
}

/*
D2I - convert stack top double value to int
*/
type D2I struct {
	base.NoOperandInstruction
}

func (inst *D2I) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	stack.PushInt(int32(d))
}

/*
D2L - convert stack top double value to long
*/
type D2L struct {
	base.NoOperandInstruction
}

func (inst *D2L) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	stack.PushLong(int64(d))
}
