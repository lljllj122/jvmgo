package base

import "jvmgo/ch06/rtda"

func Branch(frame *rtda.StackFrame, offset int) {
	pc := frame.Thread().Pc()
	nextPc := pc + offset
	frame.SetNextPc(nextPc)
}
