package classfile

/*
“LineNumberTable_attribute {
   u2 attribute_name_index;
   u4 attribute_length;
   u2 line_number_table_length;
   {   u2 start_pc;
       u2 line_number;
   } line_number_table[line_number_table_length];
}”

store information of line numbers of methods
*/
type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    u2
	lineNumber u2
}

func (attr *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	attr.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range attr.lineNumberTable {
		attr.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}
