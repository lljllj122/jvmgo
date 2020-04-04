package base

import "jvmgo/ch05/rtda"

type Instruction interface {
	// Load operand from the bytecode
	FetchOperands(reader *BytecodeReader)
	// Execute the instruction in stack frame
	Execute(frame *rtda.StackFrame)
}

type NoOperandInstruction struct {
}

func (ins *NoOperandInstruction) FetchOperands(reader *BytecodeReader) {
	// Nothing
}

/*
BranchInstruction - A branch is an instruction in a computer program that can cause a
computer to begin executing a different instruction sequence and thus deviate from its
default behavior of executing instructions in order.
*/
type BranchInstruction struct {
	Offset int16
}

func (ins *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	ins.Offset = int16(reader.ReadUint32())
}

// Index8Instruction - used to load local variable by index
type Index8Instruction struct {
	Index uint8
}

func (ins *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	ins.Index = uint8(reader.ReadByte())
}

/*
Index16Instruction - used to load operand from Constant pool.
constant_pool_count is represented by 2 bytes
*/
type Index16Instruction struct {
	Index uint16
}

func (ins *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	ins.Index = uint16(reader.ReadUint16())
}
