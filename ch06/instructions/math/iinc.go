package math

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

/*
IINC - Increment local variable by a constant
No change on OpStack
*/
type IINC struct {
	// an index into the local variable array of the current frame
	Index uint
	// an immediate signed byte, it will be sign-extended to an int
	Const int32
}

func (inst *IINC) FetchOperands(reader *base.BytecodeReader) {
	inst.Index = uint(reader.ReadByte())
	inst.Const = int32(reader.ReadByte())
}

func (inst *IINC) Execute(frame *rtda.StackFrame) {
	localVars := frame.LocalVars()
	// The local variable at index must contain an int.
	val := localVars.GetInt(inst.Index)
	// value const is first sign-extended to an int
	// and then the local variable at index is incremented by that amount.
	val = val + inst.Const
	localVars.SetInt(inst.Index, val)
}
