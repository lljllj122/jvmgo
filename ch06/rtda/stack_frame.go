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

func NewStackFrame(thread *Thread, maxLocals, maxStack uint) *StackFrame {
	return &StackFrame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
		thread:       thread,
	}
}

func (frame *StackFrame) Thread() *Thread {
	return frame.thread
}

func (frame *StackFrame) OperandStack() *OperandStack {
	return frame.operandStack
}

func (frame *StackFrame) LocalVars() LocalVars {
	return frame.localVars
}

func (frame *StackFrame) NextPc() int {
	return frame.nextPc
}

func (frame *StackFrame) SetNextPc(nextPc int) {
	frame.nextPc = nextPc
}
