package constants

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type ACONST_NULL struct {
	base.NoOperandInstruction
}

func (inst *ACONST_NULL) Execute(frame *rtda.StackFrame) {
	frame.OperandStack().PushRef(nil)
}

type DCONST_0 struct {
	base.NoOperandInstruction
}

func (inst *DCONST_0) Execute(frame *rtda.StackFrame) {
	frame.OperandStack().PushDouble(0.0)
}

type DCONST_1 struct {
	base.NoOperandInstruction
}

func (inst *DCONST_1) Execute(frame *rtda.StackFrame) {
	frame.OperandStack().PushDouble(1.0)
}

type FCONST_0 struct {
	base.NoOperandInstruction
}

func (inst *FCONST_0) Execute(frame *rtda.StackFrame) {
	frame.OperandStack().PushDouble(0.0)
}

type FCONST_1 struct {
	base.NoOperandInstruction
}

func (inst *FCONST_1) Execute(frame *rtda.StackFrame) {
	frame.OperandStack().PushDouble(1.0)
}

type FCONST_2 struct {
	base.NoOperandInstruction
}

func (inst *FCONST_2) Execute(frame *rtda.StackFrame) {
	frame.OperandStack().PushDouble(2.0)
}

type ICONST_M1 struct {
	base.NoOperandInstruction
}

func (inst *ICONST_M1) Execute(frame *rtda.StackFrame) {
	frame.OperandStack().PushInt(-1)
}

type ICONST_0 struct {
	base.NoOperandInstruction
}

func (inst *ICONST_0) Execute(frame *rtda.StackFrame) {
	frame.OperandStack().PushInt(0)
}

type ICONST_1 struct {
	base.NoOperandInstruction
}

func (inst *ICONST_1) Execute(frame *rtda.StackFrame) {
	frame.OperandStack().PushInt(1)
}

type ICONST_2 struct {
	base.NoOperandInstruction
}

func (inst *ICONST_2) Execute(frame *rtda.StackFrame) {
	frame.OperandStack().PushInt(2)
}

type ICONST_3 struct {
	base.NoOperandInstruction
}

func (inst *ICONST_3) Execute(frame *rtda.StackFrame) {
	frame.OperandStack().PushInt(3)
}

type ICONST_4 struct {
	base.NoOperandInstruction
}

func (inst *ICONST_4) Execute(frame *rtda.StackFrame) {
	frame.OperandStack().PushInt(4)
}

type ICONST_5 struct {
	base.NoOperandInstruction
}

func (inst *ICONST_5) Execute(frame *rtda.StackFrame) {
	frame.OperandStack().PushInt(5)
}

type LCONST_0 struct {
	base.NoOperandInstruction
}

func (inst *LCONST_0) Execute(frame *rtda.StackFrame) {
	frame.OperandStack().PushLong(0)
}

type LCONST_1 struct {
	base.NoOperandInstruction
}

func (inst *LCONST_1) Execute(frame *rtda.StackFrame) {
	frame.OperandStack().PushLong(1)
}
