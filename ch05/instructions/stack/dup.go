package stack

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

/*
DUP - Duplicate the value at top of Opstack and push

OpStack:
..., value →
..., value, value
*/
type DUP struct {
	base.NoOperandInstruction
}

func (inst *DUP) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	// we pass by value here, so the slot will be physically duplicated
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

/*
DUP_X1 - Duplicate the top operand stack value and insert two values down

OpStack:
..., value2, value1 →
..., value1, value2, value1
*/
type DUP_X1 struct {
	base.NoOperandInstruction
}

func (inst *DUP_X1) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}
