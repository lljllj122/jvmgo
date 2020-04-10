package stores

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// Get Long variable from opstack and Store into local variable
func _lstore(frame *rtda.StackFrame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

type LSTORE struct {
	base.Index8Instruction
}

func (inst *LSTORE) Execute(frame *rtda.StackFrame) {
	_lstore(frame, inst.Index)
}

type LSTORE_0 struct {
	base.NoOperandInstruction
}

func (inst *LSTORE_0) Execute(frame *rtda.StackFrame) {
	_lstore(frame, 0)
}

type LSTORE_1 struct {
	base.NoOperandInstruction
}

func (inst *LSTORE_1) Execute(frame *rtda.StackFrame) {
	_lstore(frame, 1)
}

type LSTORE_2 struct {
	base.NoOperandInstruction
}

func (inst *LSTORE_2) Execute(frame *rtda.StackFrame) {
	_lstore(frame, 2)
}

type LSTORE_3 struct {
	base.NoOperandInstruction
}

func (inst *LSTORE_3) Execute(frame *rtda.StackFrame) {
	_lstore(frame, 3)
}
