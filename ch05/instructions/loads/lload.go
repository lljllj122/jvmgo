package loads

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// Load Long variable from LocalVar table and push to opstack
func _lload(frame *rtda.StackFrame, index uint8) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}

type LLOAD struct {
	base.Index8Instruction
}

func (inst *LLOAD) Execute(frame *rtda.StackFrame) {
	_lload(frame, inst.Index)
}

type LLOAD_0 struct {
	base.NoOperandInstruction
}

func (inst *LLOAD_0) Execute(frame *rtda.StackFrame) {
	_lload(frame, 0)
}

type LLOAD_1 struct {
	base.NoOperandInstruction
}

func (inst *LLOAD_1) Execute(frame *rtda.StackFrame) {
	_lload(frame, 1)
}

type LLOAD_2 struct {
	base.NoOperandInstruction
}

func (inst *LLOAD_2) Execute(frame *rtda.StackFrame) {
	_lload(frame, 2)
}

type LLOAD_3 struct {
	base.NoOperandInstruction
}

func (inst *LLOAD_3) Execute(frame *rtda.StackFrame) {
	_lload(frame, 3)
}
