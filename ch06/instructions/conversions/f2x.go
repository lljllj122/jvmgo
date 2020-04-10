package conversions

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

// Convert float to double
type F2D struct {
	base.NoOperandInstruction
}

func (inst *F2D) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	d := float64(f)
	stack.PushDouble(d)
}

// Convert float to int
type F2I struct {
	base.NoOperandInstruction
}

func (inst *F2I) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	i := int32(f)
	stack.PushInt(i)
}

// Convert float to long
type F2L struct {
	base.NoOperandInstruction
}

func (inst *F2L) Execute(frame *rtda.StackFrame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	l := int64(f)
	stack.PushLong(l)
}
