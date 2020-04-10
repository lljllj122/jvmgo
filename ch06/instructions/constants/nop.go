package constants

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type NOP struct {
	base.NoOperandInstruction
}

func (nop *NOP) Execute(frame *rtda.StackFrame) {
	// do nothing
}
