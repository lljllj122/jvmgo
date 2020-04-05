package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

/*
IREM - Remainder int
..., value1, value2 →
..., value1 % value2
*/
type IREM struct {
	base.NoOperandInstruction
}

func (inst *IREM) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero.")
	}
	v1 := stack.PopInt()
	res := v1 - (v1/v2)*v2
	stack.PushInt(res)
}

/*
LREM - Remainder long
..., value1, value2 →
..., value1 % value2
*/
type LREM struct {
	base.NoOperandInstruction
}

func (inst *LREM) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero.")
	}
	v1 := stack.PopLong()
	res := v1 - (v1/v2)*v2
	stack.PushLong(res)
}

type FREM struct {
	base.NoOperandInstruction
}

type DREM struct {
	base.NoOperandInstruction
}
