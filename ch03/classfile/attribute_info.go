package classfile

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, constantPool *ConstantPool) []AttributeInfo {
	attributeCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributeCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, constantPool)
	}
	return attributes
}

func readAttribute(reader *ClassReader, constantPool *ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrName := constantPool.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attributeInfo := newAtributeInfo(attrName, attrLen, constantPool)
	return attributeInfo
}

func newAtributeInfo(attrName string, attrLen u4, constantPool *ConstantPool) AttributeInfo {
	//TODO: switch on attrName
	return nil
}

/*
“attribute_info {
   u2 attribute_name_index;
   u4 attribute_length;
   u1 info[attribute_length];
}”
*/
type UnparsedAttribute struct {
	name   string
	length u4
	info   []byte
}

func (attr *UnparsedAttribute) readInfo(reader *ClassReader) {
	attr.info = reader.readBytes(attr.length)
}
