package heap

import (
	"fmt"
	"jvmgo/ch06/classfile"
	"jvmgo/ch06/classpath"
)

type ClassLoader struct {
	classPath   *classpath.ClassPath
	loadedClass map[string]*Class
}

func NewClassLoader(classpath *classpath.ClassPath) *ClassLoader {
	return &ClassLoader{
		classPath:   classpath,
		loadedClass: map[string]*Class{},
	}
}

func (loader *ClassLoader) LoadClass(name string) *Class {
	if class, ok := loader.loadedClass[name]; ok {
		return class
	}
	return loader.loadNonArrayClass(name)
}

func (loader *ClassLoader) readClassData(name string) ([]byte, classpath.Entry) {
	data, entry, err := loader.classPath.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

func (loader *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := loader.readClassData(name)
	class := loader.parseClass(data)
	link(class)
	fmt.Printf("[Loaded %s from %s\n]", name, entry)
	return class
}

func (loader *ClassLoader) parseClass(data []byte) *Class {
	classFile, err := classfile.ParseClassFile(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	class := newClass(classFile)
	class.loader = loader
	resolveSuperClass(class)
	resolveInterface(class)
	loader.loadedClass[class.name] = class
	return class
}

func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

func resolveInterface(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

func link(class *Class) {
	verify(class)
	prepare(class)
}

func verify(class *Class) {

}

func prepare(class *Class) {

}

func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)

}