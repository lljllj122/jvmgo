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
}

func NewStackFrame(maxLocals, maxStack uint) *StackFrame {
	return &StackFrame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}
