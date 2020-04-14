package heap

import "jvmgo/ch06/classfile"

type Field struct {
	ClassMember
	constValueIndex uint
	slotId          uint
}

func newFields(class *Class, classFileFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(classFileFields))

	for i, cff := range classFileFields {
		fields[i] = new(Field)
		fields[i].class = class
		fields[i].loadMemberInfo(cff)
		fields[i].loadAttributes(cff)
	}
	return fields
}

func (f *Field) isLongOrDouble() bool {
	return f.descriptor == "J" || f.descriptor == "D"
}

func (f *Field) loadAttributes(classFileField *classfile.MemberInfo) {
	if constValAttr := classFileField.ConstantValueAttribute(); constValAttr != nil {
		f.constValueIndex = uint(constValAttr.ConstantValueIndex())
	}
}