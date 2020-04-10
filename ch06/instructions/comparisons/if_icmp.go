package comparisons

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

/*
Pop 2 values from OpStack and compare.
Branch if int comparison succeeds
*/
func _executeBranch2i(frame *rtda.StackFrame, offset int, condition func(v1 int32, v2 int32) bool) {
	v2 := frame.OperandStack().PopInt()
	v1 := frame.OperandStack().PopInt()
	if condition(v1, v2) {
		base.Branch(frame, offset)
	}
}

// v1 == v2
type IF_ICMPEQ struct {
	base.BranchInstruction
}

func (inst *IF_ICMPEQ) Execute(frame *rtda.StackFrame) {
	_executeBranch2i(frame, inst.Offset, func(v1 int32, v2 int32) bool {
		return v1 == v2
	})
}

// v1 != v2
type IF_ICMPNE struct {
	base.BranchInstruction
}

func (inst *IF_ICMPNE) Execute(frame *rtda.StackFrame) {
	_executeBranch2i(frame, inst.Offset, func(v1 int32, v2 int32) bool {
		return v1 != v2
	})
}

// v1 < v2
type IF_ICMPLT struct {
	base.BranchInstruction
}

func (inst *IF_ICMPLT) Execute(frame *rtda.StackFrame) {
	_executeBranch2i(frame, inst.Offset, func(v1 int32, v2 int32) bool {
		return v1 < v2
	})
}

// v1 <= v2
type IF_ICMPLE struct {
	base.BranchInstruction
}

func (inst *IF_ICMPLE) Execute(frame *rtda.StackFrame) {
	_executeBranch2i(frame, inst.Offset, func(v1 int32, v2 int32) bool {
		return v1 <= v2
	})
}

// v1 > v2
type IF_ICMPGT struct {
	base.BranchInstruction
}

func (inst *IF_ICMPGT) Execute(frame *rtda.StackFrame) {
	_executeBranch2i(frame, inst.Offset, func(v1 int32, v2 int32) bool {
		return v1 > v2
	})
}

// v1 >= v2
type IF_ICMPGE struct {
	base.BranchInstruction
}

func (inst *IF_ICMPGE) Execute(frame *rtda.StackFrame) {
	_executeBranch2i(frame, inst.Offset, func(v1 int32, v2 int32) bool {
		return v1 > v2
	})
}
