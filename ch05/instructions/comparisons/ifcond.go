package comparisons

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// If equal
type IFEQ struct {
	base.BranchInstruction
}

func (inst *IFEQ) Execute(frame *rtda.StackFrame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, inst.Offset)
	}
}
