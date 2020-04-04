package classfile

/*
Fields, methods, and interface methods are represented by similar structures:
CONSTANT_Fieldref_info {
	u1 tag;
	u2 class_index;
	u2 name_and_type_index;
}
CONSTANT_Methodref_info {
	u1 tag;
	u2 class_index;
	u2 name_and_type_index;
}
CONSTANT_InterfaceMethodref_info {
	u1 tag;
	u2 class_index;
	u2 name_and_type_index;
}

name of the class
type & name of the method/field
*/
type ConstantMemberRefInfo struct {
	pool             *ConstantPool
	classIndex       u2
	nameAndTypeIndex u2
}

func (info *ConstantMemberRefInfo) readInfo(reader *ClassReader) {
	info.classIndex = reader.readUint16()
	info.nameAndTypeIndex = reader.readUint16()
}

func (info *ConstantMemberRefInfo) ClassName() string {
	return info.pool.getClassName(info.classIndex)
}

func (info *ConstantMemberRefInfo) NameAndDescriptor() (string, string) {
	return info.pool.getNameAndType(info.nameAndTypeIndex)
}

type ConstantFieldRefInfo struct {
	ConstantMemberRefInfo
}

type ConstantMethodRefInfo struct {
	ConstantMemberRefInfo
}

type InterfaceMethodReInfo struct {
	ConstantMemberRefInfo
}
