package classfile

import "fmt"

/*
ClassFile {
	u4 magic;
	u2 minor_version;
	u2 major_version;
	u2 constant_pool_count;
	cp_info constant_pool[constant_pool_count-1]; u2 access_flags;
	u2 this_class;
	u2 super_class;
	u2 interfaces_count;
	u2 interfaces[interfaces_count];
	u2 fields_count;
	field_info fields[fields_count];
	u2 methods_count;
	method_info methods[methods_count];
	u2 attributes_count;
	attribute_info attributes[attributes_count];
}

ClassFile - represents structure of class file
*/
type ClassFile struct {
	magic        u4
	minorVersion u2
	majorVersion u2
	accessFlags  u2
	constantPool *ConstantPool
	thisClass    u2
	superClass   u2
	interfaces   []u2
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

// ClassName : Get the name of class
func (cf *ClassFile) ClassName() string {
	// return cf.constantPool.getClassName(cf.thisClass)
	return ""
}

// SuperClassName : Get the super class name
func (cf *ClassFile) SuperClassName() string {
	// if cf.superClass > 0 {
	// 	return cf.constantPool.getClassName(cf.superClass)
	// }
	return ""
}

// InterfaceNames : Get the interface names
func (cf *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(cf.interfaces))
	// for i, interface := range cf.interfaces {
	// 	inteinterfaceNames[i] = cf.constantPool.getClassName(interface)
	// }
	return interfaceNames
}

func (cf *ClassFile) read(reader *ClassReader) {
	// init class file content
}

func (cf *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
	cf.magic = magic
}

func (cf *ClassFile) readAndCheckVersion(reader *ClassReader) {
	cf.minorVersion = reader.readUint16()
	cf.majorVersion = reader.readUint16()

	switch cf.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if cf.minorVersion == 0 {
			return
		}
	}

	panic("java.lang.UnsupportedClassVersionError")
}

// ParseClass : Parse the input byte data into ClassFile entity
func ParseClass(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	// cr := &ClassReader{classData}
	cf = &ClassFile{}
	return
}
