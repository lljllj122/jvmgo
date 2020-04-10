package loads

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

// Load Double variable from LocalVar table and push to opstack
func _dload(frame *rtda.StackFrame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}

type DLOAD struct {
	base.Index8Instruction
}

func (inst *DLOAD) Execute(frame *rtda.StackFrame) {
	_dload(frame, inst.Index)
}

type DLOAD_0 struct {
	base.NoOperandInstruction
}

func (inst *DLOAD_0) Execute(frame *rtda.StackFrame) {
	_dload(frame, 0)
}

type DLOAD_1 struct {
	base.NoOperandInstruction
}

func (inst *DLOAD_1) Execute(frame *rtda.StackFrame) {
	_dload(frame, 1)
}

type DLOAD_2 struct {
	base.NoOperandInstruction
}

func (inst *DLOAD_2) Execute(frame *rtda.StackFrame) {
	_dload(frame, 2)
}

type DLOAD_3 struct {
	base.NoOperandInstruction
}

func (inst *DLOAD_3) Execute(frame *rtda.StackFrame) {
	_dload(frame, 3)
}
