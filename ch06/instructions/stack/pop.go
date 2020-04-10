package stack

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

/*
POP* instructions pop the top values from Opstack
POP is used to pop the int, float, and Ref types
*/
type POP struct {
	base.NoOperandInstruction
}

func (inst *POP) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

/*
POP2 is used to pop the long and double types
*/
type POP2 struct {
	base.NoOperandInstruction
}

func (inst *POP2) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
