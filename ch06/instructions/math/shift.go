package math

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

/*
ISHL - Shift left int

OpStack:
..., value1, value2 →
..., result (value1 << s)

result is calculated by shifting value1 left by s bit positions
s is the value of the low 5 bits of value2.
*/
type ISHL struct {
	base.NoOperandInstruction
}

func (inst *ISHL) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	// Get the low 5 bits of v2
	s := uint(v2) & 0x1f
	res := v1 << s
	stack.PushInt(res)
}

/*
ISHR - Arithmetic Shift right int

OpStack:
..., value1, value2 →
..., result (value1 >> s)

result is calculated by shifting value1 right by s bit positions, with sign extension
s is the value of the low 5 bits of value2.
*/
type ISHR struct {
	base.NoOperandInstruction
}

func (inst *ISHR) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	// Get the low 5 bits of v2
	s := uint(v2) & 0x1f
	res := v1 >> s
	stack.PushInt(res)
}

/*
IUSHR - Logical shift right int

OpStack:
..., value1, value2 →
..., result (value1 >>> s)

result is calculated by shifting value1 right by s bit positions, with zero extension
s is the value of the low 5 bits of value2.
*/
type IUSHR struct {
	base.NoOperandInstruction
}

func (inst *IUSHR) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	// Get the low 5 bits of v2
	s := uint(v2) & 0x1f
	// Golang doesn't support >>> so we convert the value to uint32
	res := int32(uint32(v1) >> s)
	stack.PushInt(res)
}

/*
LSHL - Shift left long

OpStack:
..., value1, value2 →
..., result (value1 << s)

The value1 must be of type long, and value2 must be of type int.
result is calculated by shifting value1 left by s bit positions
s is the value of the low 6 bits of value2.
*/
type LSHL struct {
	base.NoOperandInstruction
}

func (inst *LSHL) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	// value2 must be of type int.
	v1 := stack.PopLong()
	s := uint(v2) & 0x3f
	res := v1 << s
	stack.PushLong(res)
}

/*
LSHR - Shift right long

OpStack:
..., value1, value2 →
..., result (value1 >> s)

The value1 must be of type long, and value2 must be of type int.
result is calculated by shifting value1 right by s bit positions, with sign extension
s is the value of the low 6 bits of value2.
*/
type LSHR struct {
	base.NoOperandInstruction
}

func (inst *LSHR) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	// value2 must be of type int.
	v1 := stack.PopLong()
	s := uint(v2) & 0x3f
	res := v1 >> s
	stack.PushLong(res)
}

/*
LUSHR - Logical shift right long

OpStack:
..., value1, value2 →
..., result (value1 >>> s)

result is calculated by shifting value1 right logically by s bit positions, with zero extension
s is the value of the low 5 bits of value2.
*/
type LUSHR struct {
	base.NoOperandInstruction
}

func (inst *LUSHR) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	// Get the low 5 bits of v2
	s := uint(v2) & 0x3f
	// Golang doesn't support >>> so we convert the value to uint64
	res := int64(uint64(v1) >> s)
	stack.PushLong(res)
}
