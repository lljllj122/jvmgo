package heap

import "jvmgo/ch06/classfile"

/*
The runtime Class in memory
*/
type Class struct {
	accessFlags    uint16
	name           string
	superClassName string
	interfaceNames []string
	// constantPool   *ConstantPool
	// fields
	// method
	// loader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	// staticVars *Slot
}

// Create a runtime class from classfile
func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{
		accessFlags:    cf.AccessFlags(),
		name:           cf.ClassName(),
		superClassName: cf.SuperClassName(),
		interfaceNames: cf.InterfaceNames(),
		//superClass:        nil,
		//interfaces:        nil,
		//instanceSlotCount: 0,
		//staticSlotCount:   0,
	}

	return class
}

func (class *Class) IsPublic() bool {
	return 0 != class.accessFlags&ACC_PUBLIC
}

func (class *Class) IsFinal() bool {
	return 0 != class.accessFlags&ACC_FINAL
}

func (class *Class) IsSuper() bool {
	return 0 != class.accessFlags&ACC_SUPER
}

func (class *Class) IsInterface() bool {
	return 0 != class.accessFlags&ACC_INTERFACE
}

func (class *Class) IsAbstract() bool {
	return 0 != class.accessFlags&ACC_ABSTRACT
}

func (class *Class) IsSynthetic() bool {
	return 0 != class.accessFlags&ACC_SYNTHETIC
}

func (class *Class) IsAnnotation() bool {
	return 0 != class.accessFlags&ACC_ANNOTATION
}

func (class *Class) IsEnum() bool {
	return 0 != class.accessFlags&ACC_ENUM
}
