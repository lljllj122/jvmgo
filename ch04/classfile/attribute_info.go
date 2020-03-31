package classfile

const (
	CodeAttr               = "Code"
	ConstantValueAttr      = "ConstantValue"
	DeprecatedAttr         = "Deprecated"
	ExceptionAttr          = "Exceptions"
	LineNumberTableAttr    = "LineNumberTable"
	LocalVariableTableAttr = "LocalVariableTable"
	SourceFileAttr         = "SourceFile"
	SyntheticAttr          = "Synthetic"
)

/*
“attribute_info {
   u2 attribute_name_index;
   u4 attribute_length;
   u1 info[attribute_length];
}”
*/
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
	attributeInfo.readInfo(reader)
	return attributeInfo
}

func newAtributeInfo(attrName string, attrLen u4, constantPool *ConstantPool) AttributeInfo {
	var attribute AttributeInfo
	switch attrName {
	case CodeAttr:
		attribute = &CodeAttribute{constantPool: constantPool}
	case ConstantValueAttr:
		attribute = &ConstantValueAttribute{}
	case DeprecatedAttr:
		attribute = &DeprecatedAttribute{}
	case ExceptionAttr:
		attribute = &ExceptionsAttribute{}
	case LineNumberTableAttr:
		attribute = &LineNumberTableAttribute{}
	case LocalVariableTableAttr:
		attribute = &LocalVariableTableAttribute{}
	case SourceFileAttr:
		attribute = &SourceFileAttribute{constantPool: constantPool}
	case SyntheticAttr:
		attribute = &SyntheticAttribute{}
	default:
		attribute = &UnparsedAttribute{name: attrName, length: attrLen}
	}
	return attribute
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
