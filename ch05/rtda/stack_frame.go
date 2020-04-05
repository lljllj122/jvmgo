package rtda

/*
A new frame is created each time a method is invoked.
A frame is destroyed when its method invocation completes,
whether that completion is normal or abrupt
*/
type StackFrame struct {
	lowerFrame   *StackFrame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	nextPc       int
}

func NewStackFrame(thread *Thread, maxLocals, maxStack uint8) *StackFrame {
	return &StackFrame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
		thread:       thread,
	}
}

func (stack *StackFrame) Thread() *Thread {
	return stack.thread
}

func (stack *StackFrame) OperandStack() *OperandStack {
	return stack.operandStack
}

func (stack *StackFrame) LocalVars() LocalVars {
	return stack.localVars
}

func (stack *StackFrame) NextPc() int {
	return stack.nextPc
}

func (stack *StackFrame) SetNextPc(nextPc int) {
	stack.nextPc = nextPc
}
