package classfile

/*
CONSTANT_Utf8_info {
   u1 tag;
   u2 length;
   u1 bytes[length];
}
*/
type ConstantUtf8Info struct {
	str string
}

func (info *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	info.str = decodeMUTF8(bytes)
}

// Java is using MUTF-8 while Go is using Standard UTF-8
// TODO: This function will be revised to use actual MUTF-8
func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
