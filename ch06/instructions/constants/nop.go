package constants

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type NOP struct {
	base.NoOperandInstruction
}

func (nop *NOP) Execute(frame *rtda.StackFrame) {
	// do nothing
}
