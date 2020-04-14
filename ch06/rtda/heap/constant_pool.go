package heap

import (
	"fmt"
	"jvmgo/ch06/classfile"
)

type Constant interface {
}

type ConstantPool struct {
	class  *Class
	consts []Constant
}

func (cp *ConstantPool) GetConstant(i uint) Constant {
	if c := cp.consts[i]; c != nil {
		return c
	}

	panic(fmt.Sprintf("No constant at index %d", i))
}

func newConstantPool(class *Class, cfCp *classfile.ConstantPool) *ConstantPool {
	cpCount := cfCp.Size()
	consts := make([]Constant, cpCount)
	runtimeConstantPool := &ConstantPool{
		class:  class,
		consts: consts,
	}
	for i := uint(1); i < cpCount; i++ {
		cpInfo := runtimeConstantPool.GetConstant(i)
		switch cpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			consts[i] = intInfo.Value()
		case *classfile.ConstantFloatInfo:
			floatInfo := cpInfo.(*classfile.ConstantFloatInfo)
			consts[i] = floatInfo.Value()
		case *classfile.ConstantLongInfo:
			longInfo := cpInfo.(*classfile.ConstantLongInfo)
			consts[i] = longInfo.Value()
			i++
		case *classfile.ConstantDoubleInfo:
			doubleInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			consts[i] = doubleInfo.Value()
			i++
		case *classfile.ConstantStringInfo:
			// string is a reference constant
			stringInfo := cpInfo.(*classfile.ConstantStringInfo)
			consts[i] = stringInfo.String()
		case *classfile.ConstantClassInfo:
			classInfo := cpInfo.(*classfile.ConstantClassInfo)
			consts[i] = newClassRef(runtimeConstantPool, classInfo)
		case *classfile.ConstantFieldRefInfo:
			fieldInfo := cpInfo.(*classfile.ConstantFieldRefInfo)
			consts[i] = newFieldRef(runtimeConstantPool, fieldInfo)
		case *classfile.ConstantMethodRefInfo:
			methodInfo := cpInfo.(*classfile.ConstantMethodRefInfo)
			consts[i] = newMethodRef(runtimeConstantPool, methodInfo)
		case *classfile.ConstantInterfaceMethodReInfo:
			interfaceMethodInfo := cpInfo.(*classfile.ConstantInterfaceMethodReInfo)
			consts[i] = newInterfaceMethodRef(runtimeConstantPool, interfaceMethodInfo)
		}
	}
	return runtimeConstantPool
}
