package stores

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// Get Double variable from opstack and Store into local variable
func _dstore(frame *rtda.StackFrame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}

type DSTORE struct {
	base.Index8Instruction
}

func (inst *DSTORE) Execute(frame *rtda.StackFrame) {
	_dstore(frame, inst.Index)
}

type DSTORE_0 struct {
	base.NoOperandInstruction
}

func (inst *DSTORE_0) Execute(frame *rtda.StackFrame) {
	_dstore(frame, 0)
}

type DSTORE_1 struct {
	base.NoOperandInstruction
}

func (inst *DSTORE_1) Execute(frame *rtda.StackFrame) {
	_dstore(frame, 1)
}

type DSTORE_2 struct {
	base.NoOperandInstruction
}

func (inst *DSTORE_2) Execute(frame *rtda.StackFrame) {
	_dstore(frame, 2)
}

type DSTORE_3 struct {
	base.NoOperandInstruction
}

func (inst *DSTORE_3) Execute(frame *rtda.StackFrame) {
	_dstore(frame, 3)
}
