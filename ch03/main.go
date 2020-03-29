package main

import (
	"fmt"
	"jvmgo/ch03/classfile"
	"jvmgo/ch03/classpath"
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
	// convert the dot formatted class name to the slah formatted relative path
	className := strings.Replace(cmd.class, ".", "/", -1)
	data, _, err := classPath.ReadClass(className)

	if err != nil {
		fmt.Printf("Cannot find or load class: %v\n", cmd.class)
	}

	fmt.Printf("class data: %v\n", data)
}

func loadClass(className string, classPath *classpath.ClassPath) *classfile.ClassFile {
	classData, _, err := classPath.ReadClass(className)
	if err != nil {
		panic(err)
	}
	classFile, err := classfile.ParseClass(classData)
	if err != nil {
		panic(err)
	}
	return classFile
}

func printClassInfo(classFile *classfile.ClassFile) {
	// fmt.Printf("version: %v.%v\n", classFile.)
}
