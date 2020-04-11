package heap

import "jvmgo/ch06/classfile"

type ClassRef struct {
	SymbolRef
}

func newClassRef(cp *ConstantPool, classInfo *classfile.ConstantClassInfo) *ClassRef{
	ref := new(ClassRef)
	ref.constantPool = cp
	ref.className = classInfo.Name()
	return ref
}