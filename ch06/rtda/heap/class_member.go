package heap

import "jvmgo/ch06/classfile"

type ClassMember struct {
	accessFlag uint16
	name       string
	descriptor string
	class      *Class
}

func (member *ClassMember) loadMemberInfo(input *classfile.MemberInfo) {
	member.accessFlag = input.AccessFlag()
	member.name = input.Name()
	member.descriptor = input.Descriptor()
}
