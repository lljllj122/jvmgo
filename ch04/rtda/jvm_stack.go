package rtda

/*
Each Java Virtual Machine thread has a private Java Virtual Machine stack, created at the same time as the thread.
*/
type JvmStack struct {
	maxSize  uint
	size     uint
	topFrame *StackFrame
}

func newJvmStack(maxSize uint) *JvmStack {
	return &JvmStack{
		maxSize: maxSize,
	}
}

func (stack *JvmStack) push(frame *StackFrame) {
	if stack.size > stack.maxSize {
		panic("java.lang.StackOverFlowError")
	}

	if stack.topFrame != nil {
		frame.lowerFrame = stack.topFrame
	}
	stack.topFrame = frame
	stack.size++
}

func (stack *JvmStack) pop() *StackFrame {
	output := stack.topFrame
	stack.topFrame = output.lowerFrame
	stack.size--
	output.lowerFrame = nil
	return output
}

func (stack *JvmStack) top() *StackFrame {
	if stack.size == 0 {
		panic("jvm stack is empty.")
	}
	return stack.topFrame
}
