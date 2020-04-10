package heap

import "jvmgo/ch06/classfile"

type Field struct {
	ClassMember
}

func newFields(class *Class, classFileFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(classFileFields))

	for i, cff := range classFileFields {
		fields[i] = new(Field)
		fields[i].class = class
		fields[i].loadMemberInfo(cff)
	}

	return fields
}
