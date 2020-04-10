package control

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type TABLE_SWITCH struct {
	defaultOffset int32   // default offset to jump
	low           int32   // begin of case range
	high          int32   // end of case range
	jumpOffset    []int32 // there are (high - low + 1) number of cases
}

func (inst *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	// there are 0-3 bytes act as padding to let default offset
	// begins at an address that is a multiple of four bytes from
	// the start of the current method
	reader.SkipPadding()
	inst.defaultOffset = int32(reader.ReadUint32())
	inst.low = int32(reader.ReadUint32())
	inst.high = int32(reader.ReadUint32())
	jumpOffsetCount := inst.high - inst.low + 1
	inst.jumpOffset = make([]int32, jumpOffsetCount)
	for i := range inst.jumpOffset {
		inst.jumpOffset[i] = int32(reader.ReadUint32())
	}
}

func (inst *TABLE_SWITCH) Execute(frame *rtda.StackFrame) {
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= inst.low && index <= inst.high {
		offset = int(index)
	} else {
		offset = int(inst.defaultOffset)
	}
	base.Branch(frame, offset)
}
