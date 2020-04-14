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

func (member *ClassMember) Class() *Class {
	return member.class
}

func (member *ClassMember) Descriptor() string {
	return member.descriptor
}

func (member *ClassMember) Name() string {
	return member.name
}

func (member *ClassMember) IsPublic() bool {
	return 0 != member.accessFlag&ACC_PUBLIC
}
func (member *ClassMember) IsPrivate() bool {
	return 0 != member.accessFlag&ACC_PRIVATE
}
func (member *ClassMember) IsProtected() bool {
	return 0 != member.accessFlag&ACC_PROTECTED
}
func (member *ClassMember) IsStatic() bool {
	return 0 != member.accessFlag&ACC_STATIC
}
func (member *ClassMember) IsFinal() bool {
	return 0 != member.accessFlag&ACC_FINAL
}
func (member *ClassMember) IsSynthetic() bool {
	return 0 != member.accessFlag&ACC_SYNTHETIC
}
