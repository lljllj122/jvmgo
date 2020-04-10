package classfile

/*
“Exceptions_attribute {
   u2 attribute_name_index;
   u4 attribute_length;
   u2 number_of_exceptions;
   u2 exception_index_table[number_of_exceptions];
}”
Long attribute. Used to represent exceptions throwed by method.
*/
type ExceptionsAttribute struct {
	exceptionIndexTable []u2
}

func (attr *ExceptionsAttribute) readInfo(reader *ClassReader) {
	attr.exceptionIndexTable = reader.readUint16Table()
}

func (attr *ExceptionsAttribute) ExceptionIndexTable() []u2 {
	return attr.exceptionIndexTable
}
