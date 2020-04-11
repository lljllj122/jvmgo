package heap

import "jvmgo/ch06/classfile"

type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(constantPool *ConstantPool, refInfo *classfile.ConstantFieldRefInfo) *FieldRef {
	ref := new(FieldRef)
	ref.constantPool = constantPool
	ref.loadMemberRef(&refInfo.ConstantMemberRefInfo)
	return ref
}
