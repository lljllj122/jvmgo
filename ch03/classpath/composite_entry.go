package classpath

import (
	"errors"
	"strings"
)

// CompositeEntry is basically a list of various types of entries
type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newDirEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (c CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range c {
		data, form, err := entry.readClass(className)
		// if found the target class, return it
		if err == nil {
			return data, form, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (c CompositeEntry) String() string {
	strs := []string{}
	for _, entry := range c {
		strs = append(strs, entry.String())
	}
	return strings.Join(strs, pathListSeparator)
}
