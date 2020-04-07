package main

import (
	"fmt"
	"jvmgo/ch05/classfile"
	"jvmgo/ch05/instructions"
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

func interpret(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	byteCode := codeAttr.Code()

	thread := rtda.NewThread()
	frame := rtda.NewStackFrame(thread, maxLocals, maxStack)
	thread.PushFrame(frame)

	// catch the err when panic happens
	defer catchErr(frame)
	loop(thread, byteCode)
}

func catchErr(frame *rtda.StackFrame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars: %v\n", frame.LocalVars())
		fmt.Printf("OpStack: %v\n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtda.Thread, byteCode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		pc := frame.NextPc()
		thread.SetPc(pc)

		// decode the instruction
		reader.Reset(byteCode, pc)
		opcode := reader.ReadByte()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)

		fmt.Printf("%T\n", inst)

		frame.SetNextPc(reader.Pc())

		// execute
		inst.Execute(frame)
	}
}