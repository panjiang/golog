package log

import (
	"fmt"
	"path"
	"strings"
	"time"
)

type Filename struct {
	FullName string // eg: ./logs/test.log
	Path     string // eg: ./logs
	Name     string // eg: test.log
	Ext      string // eg: log
	NamePart string // eg: test
}

func NewFilename(filename string) *Filename {
	dir, fileName := path.Split(filename)
	fileNameArr := strings.Split(fileName, ".")
	namePart := strings.Join(fileNameArr[:len(fileNameArr)-1], ".")
	fileExt := fileNameArr[len(fileNameArr)-1]
	fn := &Filename{
		FullName: filename,
		Path:     dir,
		Name:     fileName,
		Ext:      fileExt,
		NamePart: namePart,
	}
	return fn
}

func (fn *Filename) SuffixFullName(suffix string) string {
	name := fmt.Sprintf("%s_%s.%s", fn.NamePart, suffix, fn.Ext)
	return path.Join(fn.Path, name)
}

// Get the Date num string
func GetDateNumStr(t time.Time) string {
	return t.Format("20060102")
}

// Get the Date Hour num string
func GetDateHourStr(t time.Time) string {
	return t.Format("2006010215")
}
