package stores

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// Get Int variable from opstack and Store into local variable
func _istore(frame *rtda.StackFrame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}

type ISTORE struct {
	base.Index8Instruction
}

func (inst *ISTORE) Execute(frame *rtda.StackFrame) {
	_istore(frame, inst.Index)
}

type ISTORE_0 struct {
	base.NoOperandInstruction
}

func (inst *ISTORE_0) Execute(frame *rtda.StackFrame) {
	_istore(frame, 0)
}

type ISTORE_1 struct {
	base.NoOperandInstruction
}

func (inst *ISTORE_1) Execute(frame *rtda.StackFrame) {
	_istore(frame, 1)
}

type ISTORE_2 struct {
	base.NoOperandInstruction
}

func (inst *ISTORE_2) Execute(frame *rtda.StackFrame) {
	_istore(frame, 2)
}

type ISTORE_3 struct {
	base.NoOperandInstruction
}

func (inst *ISTORE_3) Execute(frame *rtda.StackFrame) {
	_istore(frame, 3)
}
