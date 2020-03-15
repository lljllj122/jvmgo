package classfile

type ConstantInfo interface {
}

type ConstantUtf8Info struct {
	tag    u1
	length u2
	bytes  []byte
	str    string
}

type ConstantLongInfo struct {
}

type ConstantDoubleInfo struct {
}

type ConstantClassInfo struct {
	tag       u1
	nameIndex u2
}

// ConstantNameAndTypeInfo is composite of name and descriptor indecies in constant pool
type ConstantNameAndTypeInfo struct {
	tag             u1
	nameIndex       u2
	descriptorIndex u2
}

func readConstantInfo(reader *ClassReader) ConstantInfo {
	return nil
}
