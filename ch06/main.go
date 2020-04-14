package main

import (
	"fmt"
	"jvmgo/ch06/classfile"
	"jvmgo/ch06/classpath"
	"strings"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("Version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJvm(cmd)
	}
}

func startJvm(cmd *Cmd) {
	classPath := classpath.Parse(cmd.XjreOption, cmd.classPathOption)
	fmt.Printf("class:%s	classpath:%s	args:%v\n", cmd.class, cmd.classPathOption, cmd.args)
	// convert the dot formatted class name to the slash formatted relative path
	className := strings.Replace(cmd.class, ".", "/", -1)
	classFile := loadClass(className, classPath)
	printClassInfo(classFile)

	mainMethod := getMainMethod(classFile)
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class :%v\n", cmd.class)
	}
}

func loadClass(className string, classPath *classpath.ClassPath) *classfile.ClassFile {
	classData, _, err := classPath.ReadClass(className)
	if err != nil {
		panic(err)
	}
	classFile, err := classfile.ParseClassFile(classData)
	if err != nil {
		panic(err)
	}
	return classFile
}

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	for _, methodInfo := range cf.Methods() {
		if methodInfo.Name() == "main" && methodInfo.Descriptor() == "([Ljava/lang/String;)V" {
			return methodInfo
		}
	}
	return nil
}

func printClassInfo(classFile *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", classFile.MajorVersion(), classFile.MinorVersion())
	fmt.Printf("constants count: %v\n", classFile.ConstantPool().Size())
	fmt.Printf("access flags: %v\n", classFile.AccessFlag())
	fmt.Printf("this class: %v\n", classFile.ClassName())
	fmt.Printf("super class: %v\n", classFile.SuperClassName())
	fmt.Printf("interfaces: %v\n", classFile.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(classFile.Fields()))
	for _, f := range classFile.Fields() {
		fmt.Println(f.Name())
	}

	fmt.Printf("methods count: %v\n", len(classFile.Methods()))
	for _, m := range classFile.Methods() {
		fmt.Println(m.Name())
	}

	fmt.Printf("attributes count: %v\n", len(classFile.Attributes()))
	for _, attr := range classFile.Attributes() {
		fmt.Printf("%T\n", attr)
	}
}
