package stack

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

/*
DUP - Duplicate the value at top of Opstack

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

/*
DUP_X2 - Duplicate the top operand stack value and insert two or three values down
depends on the compuational types of next values
OpStack:
(v1, v2, v3 are all category 1 type)
..., value3, value2, value1 →
..., value1, value3, value2, value1
OR
(v1 is category 1 type, v2 is category 2 type)
..., value2, value1 →
..., value1, value2, value1
*/
type DUP_X2 struct {
	base.NoOperandInstruction
}

func (inst *DUP_X2) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

/*
DUP2 - Duplicate the top one or two operand stack values
depends on the computational type of the top value.
Opstack:
(v1, v2 are all category 1 type)
..., value2, value1 →
..., value2, value1, value2, value1
OR
(v is category 2 type)
..., value →
..., value, value
*/
type DUP2 struct {
	base.NoOperandInstruction
}

func (ints *DUP2) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

/*
DUP2_X1 - Duplicate the top one or two operand stack values and insert two or three values down.
depends on the computational types of top values
Opstack:
(v1, v2, v3 are all category 1 type)
..., value3, value2, value1 →
..., value2, value1, value3, value2, value1
OR
(v1 is category 2 type while v2 is category 1 type)
..., value2, value1 →
..., value1, value2, value1
*/
type DUP2_X1 struct {
	base.NoOperandInstruction
}

func (inst *DUP2_X1) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

/*
DUP2_X2 -
(v1, v2, v3, v4 are all category 1 type)
..., value4, value3, value2, value1 →
..., value2, value1, value4, value3, value2, value1
OR
(v1 is category 2 type while v2 and v3 are category 1 type )
..., value3, value2, value1 →
..., value1, value3, value2, value1
*/
type DUP2_X2 struct {
	base.NoOperandInstruction
}

func (inst *DUP2_X2) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	slot4 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot4)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}
