package stores

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// Get Float variable from opstack and Store into local variable
func _fstore(frame *rtda.StackFrame, index uint8) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}

type FSTORE struct {
	base.Index8Instruction
}

func (inst *FSTORE) Store(frame *rtda.StackFrame) {
	_fstore(frame, inst.Index)
}

type FSTORE_0 struct {
	base.NoOperandInstruction
}

func (inst *FSTORE_0) Store(frame *rtda.StackFrame) {
	_fstore(frame, 0)
}

type FSTORE_1 struct {
	base.NoOperandInstruction
}

func (inst *FSTORE_1) Store(frame *rtda.StackFrame) {
	_fstore(frame, 1)
}

type FSTORE_2 struct {
	base.NoOperandInstruction
}

func (inst *FSTORE_2) Store(frame *rtda.StackFrame) {
	_fstore(frame, 2)
}

type FSTORE_3 struct {
	base.NoOperandInstruction
}

func (inst *FSTORE_3) Store(frame *rtda.StackFrame) {
	_fstore(frame, 3)
}
