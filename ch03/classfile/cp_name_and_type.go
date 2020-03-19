package classfile

/*
“CONSTANT_NameAndType_info {
   u1 tag;
   u2 name_index;
   u2 descriptor_index;
}”
ConstantNameAndTypeInfo is composite of name and descriptor indecies in constant pool
*/
type ConstantNameAndTypeInfo struct {
	tag             u1
	nameIndex       u2
	descriptorIndex u2
}

func (info *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	panic("not implemented") // TODO: Implement
}
