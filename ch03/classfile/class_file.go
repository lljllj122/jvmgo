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
	magic             u4
	minorVersion      u2
	majorVersion      u2
	accessFlags       u2
	constantPool      *ConstantPool
	thisClassIndex    u2
	superClassIndex   u2
	interfaceIndecies []u2
	fields            []*MemberInfo
	methods           []*MemberInfo
	attributes        []AttributeInfo
}

func (cf *ClassFile) read(reader *ClassReader) {
	// init class file content
	cf.readAndCheckMagic(reader)
	cf.readAndCheckVersion(reader)
	cf.constantPool = readConstantPool(reader)
	cf.accessFlags = reader.readUint16()
	cf.thisClassIndex = reader.readUint16()
	cf.superClassIndex = reader.readUint16()
	cf.interfaceIndecies = reader.readUint16Table()
	cf.fields = readMembers(reader, cf.constantPool)
	cf.methods = readMembers(reader, cf.constantPool)
	cf.attributes = readAttributes(reader, cf.constantPool)
}

func (cf *ClassFile) MinorVersion() u2 {
	return cf.minorVersion
}

func (cf *ClassFile) MajorVersion() u2 {
	return cf.majorVersion
}

func (cf *ClassFile) ConstantPool() *ConstantPool {
	return cf.constantPool
}

func (cf *ClassFile) AccessFlags() u2 {
	return cf.accessFlags
}

// ClassName : Get the name of class
func (cf *ClassFile) ClassName() string {
	return cf.constantPool.getClassName(cf.thisClassIndex)
}

// SuperClassName : Get the super class name
func (cf *ClassFile) SuperClassName() string {
	if cf.superClassIndex > 0 {
		return cf.constantPool.getClassName(cf.superClassIndex)
	}
	return ""
}

// InterfaceNames : Get the interface names
func (cf *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(cf.interfaceIndecies))
	for i, interfaceIndex := range cf.interfaceIndecies {
		interfaceNames[i] = cf.constantPool.getClassName(interfaceIndex)
	}
	return interfaceNames
}

func (cf *ClassFile) Fields() []*MemberInfo {
	return cf.fields
}

func (cf *ClassFile) Methods() []*MemberInfo {
	return cf.methods
}

func (cf *ClassFile) Attributes() []AttributeInfo {
	return cf.Attributes()
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
