package comparisons

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
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
type IF_ICOMPEQ struct {
	base.BranchInstruction
}

func (inst *IF_ICOMPEQ) Execute(frame *rtda.StackFrame) {
	_executeBranch2i(frame, inst.Offset, func(v1 int32, v2 int32) bool {
		return v1 == v2
	})
}

// v1 != v2
type IF_ICOMPNE struct {
	base.BranchInstruction
}

func (inst *IF_ICOMPNE) Execute(frame *rtda.StackFrame) {
	_executeBranch2i(frame, inst.Offset, func(v1 int32, v2 int32) bool {
		return v1 != v2
	})
}

// v1 < v2
type IF_ICOMPLT struct {
	base.BranchInstruction
}

func (inst *IF_ICOMPLT) Execute(frame *rtda.StackFrame) {
	_executeBranch2i(frame, inst.Offset, func(v1 int32, v2 int32) bool {
		return v1 < v2
	})
}

// v1 <= v2
type IF_ICOMPLE struct {
	base.BranchInstruction
}

func (inst *IF_ICOMPLE) Execute(frame *rtda.StackFrame) {
	_executeBranch2i(frame, inst.Offset, func(v1 int32, v2 int32) bool {
		return v1 <= v2
	})
}

// v1 > v2
type IF_ICOMPGT struct {
	base.BranchInstruction
}

func (inst *IF_ICOMPGT) Execute(frame *rtda.StackFrame) {
	_executeBranch2i(frame, inst.Offset, func(v1 int32, v2 int32) bool {
		return v1 > v2
	})
}

// v1 >= v2
type IF_ICOMPGE struct {
	base.BranchInstruction
}

func (inst *IF_ICOMPGE) Execute(frame *rtda.StackFrame) {
	_executeBranch2i(frame, inst.Offset, func(v1 int32, v2 int32) bool {
		return v1 > v2
	})
}
