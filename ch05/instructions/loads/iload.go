package loads

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// Load Int variable from LocalVar table and push to opstack
func _iload(frame *rtda.StackFrame, index uint8) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

type ILOAD struct {
	base.Index8Instruction
}

func (inst *ILOAD) Execute(frame *rtda.StackFrame) {
	_iload(frame, inst.Index)
}

type ILOAD_0 struct {
	base.NoOperandInstruction
}

func (inst *ILOAD_0) Execute(frame *rtda.StackFrame) {
	_iload(frame, 0)
}

type ILOAD_1 struct {
	base.NoOperandInstruction
}

func (inst *ILOAD_1) Execute(frame *rtda.StackFrame) {
	_iload(frame, 1)
}

type ILOAD_2 struct {
	base.NoOperandInstruction
}

func (inst *ILOAD_2) Execute(frame *rtda.StackFrame) {
	_iload(frame, 2)
}

type ILOAD_3 struct {
	base.NoOperandInstruction
}

func (inst *ILOAD_3) Execute(frame *rtda.StackFrame) {
	_iload(frame, 3)
}
