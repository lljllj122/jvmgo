package comparisons

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type IFNULL struct {
	base.BranchInstruction
}

func (inst *IFNULL) Execute(frame *rtda.StackFrame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, inst.Offset)
	}
}

type IFNONNULL struct {
	base.BranchInstruction
}

func (inst *IFNONNULL) Execute(frame *rtda.StackFrame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, inst.Offset)
	}
}

