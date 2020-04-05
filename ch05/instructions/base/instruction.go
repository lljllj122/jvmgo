package base

import "jvmgo/ch05/rtda"

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

Used in IF<cond> instructions
*/
type BranchInstruction struct {
	// Offset is the operand of branch instruction
	Offset int
}

func (ins *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	ins.Offset = int(reader.ReadUint16())
}

// Index8Instruction - used to load local variable by index
type Index8Instruction struct {
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
