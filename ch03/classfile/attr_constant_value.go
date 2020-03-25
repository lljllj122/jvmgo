package classfile

/*
“ConstantValue_attribute {
   u2 attribute_name_index;
   u4 attribute_length;
   u2 constantvalue_index;
}”
Only exists in field_info struct
*/
type ConstantValueAttribute struct {
	constantValueIndex u2
}

func (attr *ConstantValueAttribute) readInfo(reader *ClassReader) {
	attr.constantValueIndex = reader.readUint16()
}

func (attr *ConstantValueAttribute) ConstantValueIndex() u2 {
	return attr.constantValueIndex
}
