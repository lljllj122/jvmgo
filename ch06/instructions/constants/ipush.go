package constants

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

// BIPUSH - get a byte value from operand, sign-extended to an int value, and push to OpStack
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

// SIPUSH - get a short value from operand, sign-extended to int value, and push to OpStack
type SIPUSH struct {
	val int16
}

func (inst *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	inst.val = int16(reader.ReadUint16())
}

func (inst *SIPUSH) Execute(frame *rtda.StackFrame) {
	i := int32(inst.val)
	frame.OperandStack().PushInt(i)
}
