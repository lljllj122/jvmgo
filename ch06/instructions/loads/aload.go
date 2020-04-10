package loads

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

// Load a Reference from LocalVar table and push to opstack
func _aload(frame *rtda.StackFrame, index uint) {
	val := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(val)
}

type ALOAD struct {
	base.Index8Instruction
}

func (inst *ALOAD) Execute(frame *rtda.StackFrame) {
	_aload(frame, inst.Index)
}

type ALOAD_0 struct {
	base.NoOperandInstruction
}

func (inst *ALOAD_0) Execute(frame *rtda.StackFrame) {
	_aload(frame, 0)
}

type ALOAD_1 struct {
	base.NoOperandInstruction
}

func (inst *ALOAD_1) Execute(frame *rtda.StackFrame) {
	_aload(frame, 1)
}

type ALOAD_2 struct {
	base.NoOperandInstruction
}

func (inst *ALOAD_2) Execute(frame *rtda.StackFrame) {
	_aload(frame, 2)
}

type ALOAD_3 struct {
	base.NoOperandInstruction
}

func (inst *ALOAD_3) Execute(frame *rtda.StackFrame) {
	_aload(frame, 3)
}
