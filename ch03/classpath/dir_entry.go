package classpath

import (
	"io/ioutil"
	"path/filepath"
)

/*
DirEntry : This entry is to search class under a specific abs dir
*/
type DirEntry struct {
	// the absolute path of the diretory
	absDir string
}

func newDirEntry(path string) *DirEntry {
	asbDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir: asbDir}
}

// read class file in abs dir
func (entry *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(entry.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, entry, err
}

func (entry *DirEntry) String() string {
	return entry.absDir
}
