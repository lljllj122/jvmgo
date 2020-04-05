package comparisons

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// Compare 2 references from OpStack
func _executeBranch2a(frame *rtda.StackFrame, offset int, condition func(ref1 *rtda.Object, ref2 *rtda.Object) bool) {
	ref2 := frame.OperandStack().PopRef()
	ref1 := frame.OperandStack().PopRef()
	if condition(ref1, ref2) {
		base.Branch(frame, offset)
	}
}

// ref1 == ref2
type IF_ACMPEQ struct {
	base.BranchInstruction
}

func (inst *IF_ACMPEQ) Execute(frame *rtda.StackFrame) {
	_executeBranch2a(frame, inst.Offset, func(ref1 *rtda.Object, ref2 *rtda.Object) bool {
		return ref1 == ref2
	})
}

// ref1 != ref2
type IF_ACMPNE struct {
	base.BranchInstruction
}

func (inst *IF_ACMPNE) Execute(frame *rtda.StackFrame) {
	_executeBranch2a(frame, inst.Offset, func(ref1 *rtda.Object, ref2 *rtda.Object) bool {
		return ref1 != ref2
	})
}