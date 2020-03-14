package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildCardEntry(classPath string) CompositeEntry {
	baseDir := classPath[:len(classPath)-1] // remove *
	compositeEntry := []Entry{}

	filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// ignore other dir under baseDir
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	})
	compositeEntry = append(compositeEntry, newDirEntry(baseDir))
	return compositeEntry
}
