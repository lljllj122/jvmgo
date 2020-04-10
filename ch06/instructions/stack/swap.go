package stack

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

/*
SWAP
..., value2, value1 →
..., value1, value2
*/
type SWAP struct {
	base.NoOperandInstruction
}

func (inst *SWAP) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
