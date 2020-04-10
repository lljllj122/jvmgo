package loads

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// Load Float variable from LocalVar table and push to opstack
func _fload(frame *rtda.StackFrame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}

type FLOAD struct {
	base.Index8Instruction
}

func (inst *FLOAD) Execute(frame *rtda.StackFrame) {
	_fload(frame, inst.Index)
}

type FLOAD_0 struct {
	base.NoOperandInstruction
}

func (inst *FLOAD_0) Execute(frame *rtda.StackFrame) {
	_fload(frame, 0)
}

type FLOAD_1 struct {
	base.NoOperandInstruction
}

func (inst *FLOAD_1) Execute(frame *rtda.StackFrame) {
	_fload(frame, 1)
}

type FLOAD_2 struct {
	base.NoOperandInstruction
}

func (inst *FLOAD_2) Execute(frame *rtda.StackFrame) {
	_fload(frame, 2)
}

type FLOAD_3 struct {
	base.NoOperandInstruction
}

func (inst *FLOAD_3) Execute(frame *rtda.StackFrame) {
	_fload(frame, 3)
}
