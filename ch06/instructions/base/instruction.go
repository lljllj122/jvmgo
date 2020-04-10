package base

import "jvmgo/ch06/rtda"

type Instruction interface {
	// Load operand from the bytecode
	FetchOperands(reader *BytecodeReader)
	// Execute the instruction in stack frame
	Execute(frame *rtda.StackFrame)
}

/*
NoOperandInstruction - An instruction that executes directly on OpStack and LocalVars
*/
type NoOperandInstruction struct {
}

func (ins *NoOperandInstruction) FetchOperands(reader *BytecodeReader) {
	// Nothing
}

/*
BranchInstruction - A branch is an instruction in a computer program that can cause a
computer to begin executing a different instruction sequence and thus deviate from its
default behavior of executing instructions in order.

Used in IF_COND and IF_CMP instructions
*/
type BranchInstruction struct {
	// Offset is the operand of branch instruction
	// Most branch instructions have 16-bit offset.
	// goto_w instruction has 32-bit offset
	Offset int
}

func (ins *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	// Branch offset is read as 16-bit format
	ins.Offset = int(reader.ReadInt16())
}

// Index8Instruction - used to load local variable by index
type Index8Instruction struct {
	// use uint here because it can be extended by WIDE instruction
	Index uint
}

func (ins *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	ins.Index = uint(reader.ReadByte())
}

/*
Index16Instruction - used to load operand from Constant pool.
constant_pool_count is represented by 2 bytes
*/
type Index16Instruction struct {
	Index uint
}

func (ins *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	ins.Index = uint(reader.ReadUint16())
}
