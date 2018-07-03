package log

import (
	"fmt"
	"path"
	"strings"
	"time"
)

// Filename for resolving filename
type Filename struct {
	FullName string // eg: ./logs/test.log
	Path     string // eg: ./logs
	Name     string // eg: test.log
	Ext      string // eg: log
	NamePart string // eg: test
}

// NewFilename creates object from a filename string
func NewFilename(filename string) *Filename {
	dir, fileName := path.Split(filename)
	fileNameArr := strings.Split(fileName, ".")
	namePart := strings.Join(fileNameArr[:len(fileNameArr)-1], ".")
	fileExt := fileNameArr[len(fileNameArr)-1]
	return &Filename{
		FullName: filename,
		Path:     dir,
		Name:     fileName,
		Ext:      fileExt,
		NamePart: namePart,
	}
}

// SuffixFullName returns a full file name
func (fn *Filename) SuffixFullName(suffix string) string {
	name := fmt.Sprintf("%s_%s.%s", fn.NamePart, suffix, fn.Ext)
	return path.Join(fn.Path, name)
}

// GetDateNumStr gets the Date num string
func GetDateNumStr(t time.Time) string {
	return t.Format("20060102")
}

// GetDateHourStr gets the Date Hour num string
func GetDateHourStr(t time.Time) string {
	return t.Format("2006010215")
}
