package stores

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

// Get a Reference from opstack and Store into local variable
func _astore(frame *rtda.StackFrame, index uint) {
	val := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, val)
}

type ASTORE struct {
	base.Index8Instruction
}

func (inst *ASTORE) Execute(frame *rtda.StackFrame) {
	_astore(frame, inst.Index)
}

type ASTORE_0 struct {
	base.NoOperandInstruction
}

func (inst *ASTORE_0) Execute(frame *rtda.StackFrame) {
	_astore(frame, 0)
}

type ASTORE_1 struct {
	base.NoOperandInstruction
}

func (inst *ASTORE_1) Execute(frame *rtda.StackFrame) {
	_astore(frame, 1)
}

type ASTORE_2 struct {
	base.NoOperandInstruction
}

func (inst *ASTORE_2) Execute(frame *rtda.StackFrame) {
	_astore(frame, 2)
}

type ASTORE_3 struct {
	base.NoOperandInstruction
}

func (inst *ASTORE_3) Execute(frame *rtda.StackFrame) {
	_astore(frame, 3)
}
