package heap

import "jvmgo/ch06/classfile"

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(pool *ConstantPool, refInfo *classfile.ConstantInterfaceMethodReInfo) *InterfaceMethodRef {
	ref := new(InterfaceMethodRef)
	ref.constantPool = pool
	ref.loadMemberRef(&refInfo.ConstantMemberRefInfo)
	return ref
}