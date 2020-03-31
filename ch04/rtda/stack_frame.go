package rtda

/*
A new frame is created each time a method is invoked.
A frame is destroyed when its method invocation completes,
whether that completion is normal or abrupt
*/
type StackFrame struct {
	lower *StackFrame
	// Each frame contains an array of variables known as its local variables.
	// localVars *
	// operandStack
}

func NewStackFrame(maxLocals, maxStack uint) *StackFrame {
	return &StackFrame{
		// localVars
		// operandStack
	}
}
