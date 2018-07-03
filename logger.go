package log

// A simple log library,
// which bases on golang goverment's log module
//
// With rotate
// at certain time
// limit certain size

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Level int8

// Level enum define
const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelPanic
	LevelFatal
)

// Level's output name
const (
	NameDebug = "debug"
	NameInfo  = "info"
	NameWarn  = "warn"
	NameError = "error"
)

// Map level id with name
var levelsName = map[Level]string{
	LevelDebug: NameDebug,
	LevelInfo:  NameInfo,
	LevelWarn:  NameWarn,
	LevelError: NameError,
	LevelPanic: NameError,
	LevelFatal: NameError,
}

// Logger struct
type Logger struct {
	*log.Logger
	Name  string // Name for ouput
	Level Level  // Min level
}

// New create a new Logger
func New(name string, level Level, out io.Writer) *Logger {
	return &Logger{log.New(out, name, log.LstdFlags), name, level}
}

func (l *Logger) log(level Level, msg string) {
	if level >= l.Level {
		levelName := levelsName[level]
		var s string
		if l.Name != "" {
			s = fmt.Sprintf("%s [%s] %s", l.Name, levelName, msg)
		} else {
			s = fmt.Sprintf("[%s] %s", levelName, msg)
		}
		l.Logger.Print(s)
	}
}

func (l *Logger) Print(v ...interface{}) {
	l.log(LevelDebug, fmt.Sprint(v...))
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.log(LevelDebug, fmt.Sprintf(format, v...))
}

func (l *Logger) Println(v ...interface{}) {
	l.log(LevelDebug, fmt.Sprintln(v...))
}

func (l *Logger) Debug(v ...interface{}) {
	l.log(LevelDebug, fmt.Sprintln(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.log(LevelDebug, fmt.Sprintf(format, v...))
}

func (l *Logger) Info(v ...interface{}) {
	l.log(LevelInfo, fmt.Sprintln(v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.log(LevelInfo, fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(v ...interface{}) {
	l.log(LevelWarn, fmt.Sprintln(v...))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.log(LevelWarn, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.log(LevelError, fmt.Sprintln(v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.log(LevelError, fmt.Sprintf(format, v...))
}

func (l *Logger) Panic(v ...interface{}) {
	s := fmt.Sprintln(v...)
	l.log(LevelPanic, s)
	panic(s[:len(s)-1])
}

func (l *Logger) PanicNoPanic(v ...interface{}) string {
	s := fmt.Sprintln(v...)
	l.log(LevelPanic, s)
	fmt.Println("#######")
	return s
}

func (l *Logger) Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	panic(s)
}

func (l *Logger) PanicfNoPanic(format string, v ...interface{}) string {
	s := fmt.Sprintf(format, v...)
	l.log(LevelPanic, s)
	return s
}

func (l *Logger) Fatal(v ...interface{}) {
	l.log(LevelFatal, fmt.Sprintln(v...))
	os.Exit(1)
}

func (l *Logger) FatalNoExit(v ...interface{}) {
	l.log(LevelFatal, fmt.Sprintln(v...))
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.log(LevelFatal, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (l *Logger) FatalfNoExit(format string, v ...interface{}) {
	l.log(LevelFatal, fmt.Sprintf(format, v...))
}
