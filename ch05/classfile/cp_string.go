package classfile

/*
“CONSTANT_String_info {
   u1 tag;
   u2 string_index;
}”
it is used to represent constant objects of the type String:
it doesn't contain string content directly.
It contains an index to another UTF8_info item in constant_pool table
*/
type ConstantStringInfo struct {
	pool        *ConstantPool
	stringIndex u2
}

func (info *ConstantStringInfo) readInfo(reader *ClassReader) {
	info.stringIndex = reader.readUint16()
}

func (info *ConstantStringInfo) String() string {
	return info.pool.getUtf8(info.stringIndex)
}
