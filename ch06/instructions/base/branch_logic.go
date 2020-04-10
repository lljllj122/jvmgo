package base

import "jvmgo/ch05/rtda"

func Branch(frame *rtda.StackFrame, offset int) {
	pc := frame.Thread().Pc()
	nextPc := pc + offset
	frame.SetNextPc(nextPc)
}
