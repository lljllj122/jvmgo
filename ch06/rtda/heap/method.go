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
	return methods
}

func (m *Method) loadCodeAttribute(classFileMethod *classfile.MemberInfo) {
	if codeAttr := classFileMethod.CodeAttribute(); codeAttr != nil {
		m.maxStack = codeAttr.MaxStack()
		m.maxLocals = codeAttr.MaxLocals()
		m.byteCode = codeAttr.Code()
	}
}

func (m *Method) IsSynchronized() bool {
	return 0 != m.accessFlag&ACC_SYNCHRONIZED
}
func (m *Method) IsBridge() bool {
	return 0 != m.accessFlag&ACC_BRIDGE
}
func (m *Method) IsVarargs() bool {
	return 0 != m.accessFlag&ACC_VARARGS
}
func (m *Method) IsNative() bool {
	return 0 != m.accessFlag&ACC_NATIVE
}
func (m *Method) IsAbstract() bool {
	return 0 != m.accessFlag&ACC_ABSTRACT
}
func (m *Method) IsStrict() bool {
	return 0 != m.accessFlag&ACC_STRICT
}
