package heap

import "jvmgo/ch06/classfile"

/*
The runtime Class in memory
*/
type Class struct {
	accessFlag        uint16
	name              string
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
}

// Create a runtime class from classfile
func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{
		accessFlag:     cf.AccessFlag(),
		name:           cf.ClassName(),
		superClassName: cf.SuperClassName(),
		interfaceNames: cf.InterfaceNames(),
	}
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (class *Class) IsPublic() bool {
	return 0 != class.accessFlag&ACC_PUBLIC
}

func (class *Class) IsFinal() bool {
	return 0 != class.accessFlag&ACC_FINAL
}

func (class *Class) IsSuper() bool {
	return 0 != class.accessFlag&ACC_SUPER
}

func (class *Class) IsInterface() bool {
	return 0 != class.accessFlag&ACC_INTERFACE
}

func (class *Class) IsAbstract() bool {
	return 0 != class.accessFlag&ACC_ABSTRACT
}

func (class *Class) IsSynthetic() bool {
	return 0 != class.accessFlag&ACC_SYNTHETIC
}

func (class *Class) IsAnnotation() bool {
	return 0 != class.accessFlag&ACC_ANNOTATION
}

func (class *Class) IsEnum() bool {
	return 0 != class.accessFlag&ACC_ENUM
}
