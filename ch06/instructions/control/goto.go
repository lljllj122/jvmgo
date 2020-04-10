package control

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (inst *GOTO) Execute(frame *rtda.StackFrame) {
	base.Branch(frame, inst.Offset)
}
