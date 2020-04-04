package constants

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// BIPUSH - get a byte operand, convert to int type, and push to opstack
type BIPUSH struct {
	val byte
}

func (inst *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	inst.val = reader.ReadByte()
}

func (inst *BIPUSH) Execute(frame *rtda.StackFrame) {
	i := int32(inst.val)
	frame.OperandStack().PushInt(i)
}

// SIPUSH - get a short operand, convert to int type, and push to opstack
type SIPUSH struct {
	val int16
}

func (inst *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	inst.val = reader.ReadInt16()
}

func (inst *SIPUSH) Execute(frame *rtda.StackFrame) {
	i := int32(inst.val)
	frame.OperandStack().PushInt(i)
}
