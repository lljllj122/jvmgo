package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

/*
IINC - Increment local variable by a constant
No change on OpStack
 */
type IINC struct {
	// an index into the local variable array of the current frame
	Index uint8
	// an immediate signed byte
	Const byte
}

func (inst *IINC) FetchOperands(reader *base.BytecodeReader) {
	inst.Index = reader.ReadByte()
	inst.Const = reader.ReadByte()
}

func (inst *IINC) Execute(frame *rtda.StackFrame) {
	localVars := frame.LocalVars()
	// The local variable at index must contain an int.
	val := localVars.GetInt(inst.Index)
	// value const is first sign-extended to an int
	// and then the local variable at index is incremented by that amount.
	val = val + int32(inst.Const)
	localVars.SetInt(inst.Index, val)
}