package log

import (
	"os"
	"sync"
	"time"
)

type rotateTime uint8

// Rotate by time unit
const (
	RotateTimeClose rotateTime = iota
	RotateTimeByDay
	RotateTimeByHour
)

// Map level id with name
var rotateNames = map[string]rotateTime{
	"close": RotateTimeClose,
	"day":   RotateTimeByDay,
	"hour":  RotateTimeByHour,
}

// RotateWriter a rotate writer handler
type RotateWriter struct {
	lock          sync.Mutex
	file          *os.File
	fileName      *Filename
	rotateTimeBy  rotateTime // rotate by certain durations
	rotateTimeCur string     // store time string for rotation
}

// NewRotateWriter will create path if not exist
func NewRotateWriter(filename string, rotateTimeBy rotateTime) (*RotateWriter, error) {
	rw := &RotateWriter{
		fileName:     NewFilename(filename),
		rotateTimeBy: rotateTimeBy,
	}

	if _, err := os.Stat(rw.fileName.Path); err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(rw.fileName.Path, 0755)
		} else {
			return nil, err
		}
	}

	timeKey := rw.CurentRotateTimeKey()
	var curFileName string
	if timeKey == "" {
		curFileName = rw.fileName.FullName
	} else {
		curFileName = rw.fileName.SuffixFullName(timeKey)
	}

	rw.rotateTimeCur = timeKey
	err := rw.BindFile(curFileName)
	if err != nil {
		return nil, err
	}

	return rw, nil
}

// BindFile binds with current log file
// When startup, and rotation
func (rw *RotateWriter) BindFile(curFileName string) error {
	f, err := os.OpenFile(curFileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	rw.file = f
	return nil
}

// CurentRotateTimeKey Judges current file name
// will with date num suffix while open date rolation
func (rw *RotateWriter) CurentRotateTimeKey() string {
	switch rw.rotateTimeBy {
	case RotateTimeByDay:
		return GetDateNumStr(time.Now())
	case RotateTimeByHour:
		return GetDateHourStr(time.Now())
	default:
		return ""
	}
}

func (rw *RotateWriter) Write(p []byte) (n int, err error) {
	timeKey := rw.CurentRotateTimeKey()
	if timeKey != "" && timeKey != rw.rotateTimeCur {
		// Time unit changed
		// trigger rotation
		curFileName := rw.fileName.SuffixFullName(timeKey)
		err = rw.BindFile(curFileName)
		if err != nil {
			return 0, err
		}
	}
	n, err = rw.file.Write(p)
	if err != nil {
		return n, err
	}
	return n, err
}
