package classfile

/*
SourceFile_attribute {
   u2 attribute_name_index;
   u4 attribute_length;
   u2 sourcefile_index;
}

Only Exists in ClassFile struct
Describe the source file name
*/
type SourceFileAttribute struct {
	constantPool    *ConstantPool
	sourceFileIndex u2
}

func (attr *SourceFileAttribute) readInfo(reader *ClassReader) {
	attr.sourceFileIndex = reader.readUint16()
}

func (attr *SourceFileAttribute) FileName() string {
	return attr.constantPool.getUtf8(attr.sourceFileIndex)
}
