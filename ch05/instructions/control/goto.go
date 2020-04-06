package control

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (inst *GOTO) Execute(frame *rtda.StackFrame) {
	base.Branch(frame, inst.Offset)
}
