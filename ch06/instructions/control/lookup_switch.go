package control

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32 // (i, i+1) is a k-v pair, start from 0
}

func (inst *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	inst.defaultOffset = int32(reader.ReadUint32())
	inst.npairs = int32(reader.ReadUint32())
	inst.matchOffsets = make([]int32, inst.npairs*2)
	for i := range inst.matchOffsets {
		inst.matchOffsets[i] = int32(reader.ReadUint32())
	}
}

func (inst *LOOKUP_SWITCH) Execute(frame *rtda.StackFrame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < inst.npairs*2; i += 2 {
		// matchOffsets[i] is the key
		if inst.matchOffsets[i] == key {
			// matchOffsets[i+1] is the value (offset)
			offset := inst.matchOffsets[i+1]
			base.Branch(frame, int(offset))
			return
		}
	}
	base.Branch(frame, int(inst.defaultOffset))
}
