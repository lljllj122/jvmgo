package control

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type GOTO_W struct {
	offset int
}

func (inst *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	// offset size of GOTO_W is 32-bite
	inst.offset = int(reader.ReadUint32())
}

func (inst *GOTO_W) Execute(frame *rtda.StackFrame) {
	base.Branch(frame, inst.offset)
}

