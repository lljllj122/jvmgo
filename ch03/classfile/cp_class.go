package classfile

/*
“CONSTANT_Class_info {
   u1 tag;
   u2 name_index;
}”

represents referecen of class and interface
Class_info item contains a index to another UTF8_info item to represent the class name
*/
type ConstantClassInfo struct {
	pool      *ConstantPool
	nameIndex u2
}

func (info *ConstantClassInfo) readInfo(reader *ClassReader) {
	info.nameIndex = reader.readUint16()
}

func (info *ConstantClassInfo) Name() string {
	return info.pool.getUtf8(info.nameIndex)
}
