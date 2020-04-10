package heap

import "jvmgo/ch06/classfile"

type Method struct {
	ClassMember
	maxStack  uint
	maxLocals uint
	byteCode  []byte
}

func newMethods(class *Class, classFileMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(classFileMethods))
	for i, classFileMethod := range classFileMethods {
		methods[i] = new(Method)
		methods[i].class = class
		methods[i].loadMemberInfo(classFileMethod)
		methods[i].loadCodeAttribute(classFileMethod)
	}
}

func (m *Method) loadCodeAttribute(classFileMethod *classfile.MemberInfo) {
	if codeAttr := classFileMethod.CodeAttribute(); codeAttr != nil {
		m.maxStack = codeAttr.MaxStack()
		m.maxLocals = codeAttr.MaxLocals()
		m.byteCode = codeAttr.Code()
	}
}
