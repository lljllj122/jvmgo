package classpath

import (
	"fmt"
	"os"
	"path/filepath"
)

// ClassPath - an composite of 3 types of ClassPath: boot, ext, user
type ClassPath struct {
	bootClassPathEntry Entry
	userClassPathEntry Entry
	extClassPathEntry  Entry
}

// Parse - Parse in input from user and init ClassPath entries
func Parse(jreOption string, classpathOption string) *ClassPath {
	cp := &ClassPath{}
	cp.initBootAndExtClassPath(jreOption)
	// TODO: ext class path
	cp.initUserClassPath(classpathOption)
	return cp
}

// ReadClass - Search and load the class file from ClassPath
func (cp *ClassPath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, cpEntry, err := cp.bootClassPathEntry.readClass(className); err == nil {
		return data, cpEntry, err
	}

	if data, cpEntry, err := cp.userClassPathEntry.readClass(className); err == nil {
		return data, cpEntry, err
	}

	return cp.extClassPathEntry.readClass(className)
}

// parse the input of option -Xjre
func (cp *ClassPath) initBootAndExtClassPath(jreOption string) {
	jreDir := getJreDir(jreOption)
	// jreDir/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	fmt.Printf("Init boot class path: %v\n", jreLibPath)
	cp.bootClassPathEntry = newWildCardEntry(jreLibPath)

	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	fmt.Printf("Init ext class path: %v\n", jreExtPath)
	cp.extClassPathEntry = newWildCardEntry(jreExtPath)
}

// parse the input of option -classpath/-cp
func (cp *ClassPath) initUserClassPath(classpathOption string) {
	// if no classpath option is passed in -classpath/-cp, we will use the current dir as default
	if classpathOption != "" {
		classpathOption = "."
	}
	cp.userClassPathEntry = newDirEntry(classpathOption)
}

func (cp *ClassPath) String() string {
	return cp.userClassPathEntry.String()
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}

	if exists("./jre") {
		return "./jre"
	}

	if javaHome := os.Getenv("JAVA_HOME"); javaHome != "" {
		// $JAVA_HOME/jre
		return javaHome
	}

	panic("Cannot find a jre folder.")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
