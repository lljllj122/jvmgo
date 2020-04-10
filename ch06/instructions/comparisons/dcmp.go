package comparisons

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

/*
DCMPG - Pop top 2 double values from OpStack and compare

OpStack:
..., value1, value2 →
..., result

If value1 > value2, push 1
If value1 < value2, push -1
If value == value2, push 0

If at least one of value1 or value2 is NaN, push the int value 1 onto the operand stack
*/
type DCMPG struct {
	base.NoOperandInstruction
}

/*
If at least one of value1 or value2 is NaN, push the int value -1 onto the operand stack.
*/
type DCMPL struct {
	base.NoOperandInstruction
}

func _dcmp(frame *rtda.StackFrame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()

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

func (inst *DCMPG) Execute(frame *rtda.StackFrame) {
	_dcmp(frame, true)
}

func (inst *DCMPL) Execute(frame *rtda.StackFrame) {
	_dcmp(frame, false)
}
