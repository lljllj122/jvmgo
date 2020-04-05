package stack

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
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
