package comparisons

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

/*
LCMP - Compare top 2 long values and push a int value onto stack
*/
type LCMP struct {
	base.NoOperandInstruction
}

func (inst *LCMP) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else {
		stack.PushInt(0)
	}
}
