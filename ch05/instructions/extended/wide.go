package extended

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/instructions/loads"
	"jvmgo/ch05/instructions/math"
	"jvmgo/ch05/instructions/stores"
	"jvmgo/ch05/rtda"
)

/*
The size of local variable index for most instructions is 8-bit (unsigned).
WIDE instruction extends local variable index by additional bytes
*/
type WIDE struct {
	widenInstruction base.Instruction
}

func (inst *WIDE) FetchOperands(reader *base.BytecodeReader) {
	// the first operand of WIDE is the opcode of input instruction
	opCode := reader.ReadByte()

	switch opCode {
	case 0x15:
		widenInst := &loads.ILOAD{}
		widenInst.Index = uint(reader.ReadUint16())
		inst.widenInstruction = widenInst
	case 0x16:
		widenInst := &loads.LLOAD{}
		widenInst.Index = uint(reader.ReadUint16())
		inst.widenInstruction = widenInst
	case 0x17:
		widenInst := &loads.FLOAD{}
		widenInst.Index = uint(reader.ReadUint16())
		inst.widenInstruction = widenInst
	case 0x18:
		widenInst := &loads.DLOAD{}
		widenInst.Index = uint(reader.ReadUint16())
		inst.widenInstruction = widenInst
	case 0x19:
		widenInst := &loads.ALOAD{}
		widenInst.Index = uint(reader.ReadUint16())
		inst.widenInstruction = widenInst
	case 0x36:
		widenInst := &stores.ISTORE{}
		widenInst.Index = uint(reader.ReadUint16())
		inst.widenInstruction = widenInst
	case 0x37:
		widenInst := &stores.LSTORE{}
		widenInst.Index = uint(reader.ReadUint16())
		inst.widenInstruction = widenInst
	case 0x38:
		widenInst := &stores.FSTORE{}
		widenInst.Index = uint(reader.ReadUint16())
		inst.widenInstruction = widenInst
	case 0x39:
		widenInst := &stores.DSTORE{}
		widenInst.Index = uint(reader.ReadUint16())
		inst.widenInstruction = widenInst
	case 0x3a:
		widenInst := &stores.ASTORE{}
		widenInst.Index = uint(reader.ReadUint16())
		inst.widenInstruction = widenInst
	case 0x84:
		widenInst := &math.IINC{}
		widenInst.Index = uint(reader.ReadUint16())
		widenInst.Const = int32(reader.ReadUint16())
		inst.widenInstruction = widenInst
	case 0xa9: // ret
		panic("Unsupported opcode: 0xa9!")
	}

}

// execute the widened instruction
func (inst *WIDE) Execute(frame *rtda.StackFrame) {
	inst.widenInstruction.Execute(frame)
}
