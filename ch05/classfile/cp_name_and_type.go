package classfile

/*
“CONSTANT_NameAndType_info {
   u1 tag;
   u2 name_index;
   u2 descriptor_index;
}”
The CONSTANT_NameAndType_info structure is used to represent a field or method,
without indicating which class or interface type it belongs to.
ConstantNameAndTypeInfo is composite of name and descriptor indecies to UTF8_info items in constant pool
*/
type ConstantNameAndTypeInfo struct {
	nameIndex       u2
	descriptorIndex u2
}

func (info *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	info.nameIndex = reader.readUint16()
	info.descriptorIndex = reader.readUint16()
}
