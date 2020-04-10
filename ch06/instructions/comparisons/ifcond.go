package comparisons

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

/*
Load Operand Offset (2 bytes)
Pop x from OpStack and compare with 0
Branch if comparison succeeds
*/
func _executeBranch1i(frame *rtda.StackFrame, offset int, condition func(val int32) bool) {
	val := frame.OperandStack().PopInt()
	if condition(val) {
		base.Branch(frame, offset)
	}
}

/*
IFEQ - If Equal

if x == 0, proceeds at the offset from the address of the opcode of this instruction (branch)
*/
type IFEQ struct {
	base.BranchInstruction
}

func (inst *IFEQ) Execute(frame *rtda.StackFrame) {
	_executeBranch1i(frame, inst.Offset, func(val int32) bool {
		return val == 0
	})
}

// x != 0
type IFNE struct {
	base.BranchInstruction
}

func (inst *IFNE) Execute(frame *rtda.StackFrame) {
	_executeBranch1i(frame, inst.Offset, func(val int32) bool {
		return val != 0
	})
}

// x < 0
type IFLT struct {
	base.BranchInstruction
}

func (inst *IFLT) Execute(frame *rtda.StackFrame) {
	_executeBranch1i(frame, inst.Offset, func(val int32) bool {
		return val < 0
	})
}

// x <= 0
type IFLE struct {
	base.BranchInstruction
}

func (inst *IFLE) Execute(frame *rtda.StackFrame) {
	_executeBranch1i(frame, inst.Offset, func(val int32) bool {
		return val <= 0
	})
}

// > 0
type IFGT struct {
	base.BranchInstruction
}

func (inst *IFGT) Execute(frame *rtda.StackFrame) {
	_executeBranch1i(frame, inst.Offset, func(val int32) bool {
		return val > 0
	})
}

type IFGE struct {
	base.BranchInstruction
}

func (inst *IFGE) Execute(frame *rtda.StackFrame) {
	_executeBranch1i(frame, inst.Offset, func(val int32) bool {
		return val >= 0
	})
}
