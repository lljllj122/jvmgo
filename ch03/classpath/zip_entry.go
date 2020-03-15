package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

// ZipEntry is for loading class from .zip or .jar file
type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath: absPath}
}

func (entry *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	zipFile, err := zip.OpenReader(entry.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer zipFile.Close()

	for _, file := range zipFile.File {
		if file.Name == className {
			classFile, err := file.Open()
			if err != nil {
				return nil, nil, err
			}
			defer classFile.Close()
			data, err := ioutil.ReadAll(classFile)
			if err != nil {
				return nil, nil, err
			}
			return data, entry, nil
		}
	}

	return nil, nil, errors.New("class not found: " + className)
}

func (entry *ZipEntry) String() string {
	return entry.absPath
}
