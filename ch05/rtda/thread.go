package rtda

/*
Thread has its own pc (Program Counter) and JVM Stack
*/
type Thread struct {
	pc    int
	stack *JvmStack
}

func (thread *Thread) Pc() int {
	return thread.pc
}

func (thread *Thread) SetPc(pc int) {
	thread.pc = pc
}

func NewThread() *Thread {
	return &Thread{
		// TODO: Make the stack size configurable
		stack: newJvmStack(1024),
	}
}

func (thread *Thread) PushFrame(frame *StackFrame) {
	thread.stack.push(frame)
}

func (thread *Thread) PopFrame() *StackFrame {
	return thread.stack.pop()
}

func (thread *Thread) TopFrame() *StackFrame {
	return thread.stack.top()
}
