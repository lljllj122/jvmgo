package comparisons

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

/*
FCMPG - Pop top 2 float values from OpStack and compare

OpStack:
..., value1, value2 →
..., result

If value1 > value2, push 1
If value1 < value2, push -1
If value == value2, push 0

If at least one of value1 or value2 is NaN, push the int value 1 onto the operand stack
*/
type FCMPG struct {
	base.NoOperandInstruction
}

/*
If at least one of value1 or value2 is NaN, push the int value -1 onto the operand stack.
*/
type FCMPL struct {
	base.NoOperandInstruction
}

func _fcmp(frame *rtda.StackFrame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()

	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}

func (inst *FCMPG) Execute(frame *rtda.StackFrame) {
	_fcmp(frame, true)
}

func (inst *FCMPL) Execute(frame *rtda.StackFrame) {
	_fcmp(frame, false)
}
