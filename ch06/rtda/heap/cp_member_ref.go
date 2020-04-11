package heap

import "jvmgo/ch06/classfile"

type MemberRef struct {
	SymbolRef
	name       string
	descriptor string
}

func (ref *MemberRef) loadMemberRef(refInfo *classfile.ConstantMemberRefInfo) {
	ref.className = refInfo.ClassName()
	ref.name, ref.descriptor = refInfo.NameAndDescriptor()
}
