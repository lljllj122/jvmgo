package classfile

/*
LocalVariableTable_attribute {
	u2 attribute_name_index;
	u4 attribute_length;
	u2 local_variable_table_length;
	{ 	u2 start_pc;
		u2 length;
		u2 name_index;
		u2 descriptor_index;
		u2 index;
	} local_variable_table[local_variable_table_length]; }
*/
type LocalVariableTableAttribute struct {
	constantpool   *ConstantPool
	localVariables []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPc    u2
	length     u2
	name       string
	descriptor string
	index      u2
}

func (attr *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	localVariableTableLength := reader.readUint16()
	attr.localVariables = make([]*LocalVariableTableEntry, localVariableTableLength)
	for i := range attr.localVariables {
		attr.localVariables[i] = &LocalVariableTableEntry{
			startPc:    reader.readUint16(),
			length:     reader.readUint16(),
			name:       attr.constantpool.getUtf8(reader.readUint16()),
			descriptor: attr.constantpool.getUtf8(reader.readUint16()),
			index:      reader.readUint16(),
		}
	}
}
