package classfile

/*
“Code_attribute {
   u2 attribute_name_index;
   u4 attribute_length;
   u2 max_stack;
   u2 max_locals;
   u4 code_length;
   u1 code[code_length];
   u2 exception_table_length;
   {  u2 start_pc;
      u2 end_pc;
      u2 handler_pc;
      u2 catch_type;
   } exception_table[exception_table_length];
   u2 attributes_count;
   attribute_info attributes[attributes_count];
}”

Exists in method_info
*/
type CodeAttribute struct {
	constantPool   *ConstantPool
	maxStack       uint16     // the max depth of stack
	maxLocals      uint16     // the max size of local variables
	code           []byte // the byte code of method
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

func (attr *CodeAttribute) MaxStack() uint {
	return uint(attr.maxStack)
}

func (attr *CodeAttribute) Code() []byte {
	return attr.code
}

func (attr *CodeAttribute) ExceptionTable() []*ExceptionTableEntry {
	return attr.exceptionTable
}

func (attr *CodeAttribute) Attributes() []AttributeInfo {
	return attr.attributes
}

func (attr *CodeAttribute) MaxLocals() uint {
	return uint(attr.maxLocals)
}

func (attr *CodeAttribute) ConstantPool() *ConstantPool {
	return attr.constantPool
}

func (attr *CodeAttribute) readInfo(reader *ClassReader) {
	attr.maxStack = reader.readUint16()
	attr.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	attr.code = reader.readBytes(codeLength)
	attr.exceptionTable = readExceptionTable(reader)
	attr.attributes = readAttributes(reader, attr.constantPool)
}

type ExceptionTableEntry struct {
	startPc   u2
	endPc     u2
	handlerPc u2
	catchType u2
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}

	return exceptionTable
}