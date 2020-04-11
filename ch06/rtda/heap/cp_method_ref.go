package heap

import "jvmgo/ch06/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(constantPool *ConstantPool, refInfo *classfile.ConstantMethodRefInfo) *MethodRef {
	ref := new(MethodRef)
	ref.constantPool = constantPool
	ref.loadMemberRef(&refInfo.ConstantMemberRefInfo)
	return ref
}