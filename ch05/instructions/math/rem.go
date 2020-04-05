package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
	"math"
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

/*
FREM - Remainder float
..., value1, value2 →
..., value1 % value2
*/
type FREM struct {
	base.NoOperandInstruction
}

func (inst *FREM) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero.")
	}
	v1 := stack.PopFloat()
	res := float32(math.Mod(float64(v1), float64(v2)))
	stack.PushFloat(res)
}

/*
DREM - Remainder double
..., value1, value2 →
..., value1 % value2
*/
type DREM struct {
	base.NoOperandInstruction
}

func (inst *DREM) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero.")
	}
	v1 := stack.PopDouble()
	res := math.Mod(v1, v2)
	stack.PushDouble(res)
}
