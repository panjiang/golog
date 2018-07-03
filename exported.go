package log

import (
	"io"
	"os"
	"path"
)

var std = New("", LevelDebug, os.Stderr)
var loggers = []*Logger{}

// GetLogger fetchs this module log
func GetLogger() *Logger {
	return std
}

// AddHandler registers handler
func AddHandler(logger *Logger) {
	loggers = append(loggers, logger)
}

// SetLevel sets log output level
func SetLevel(level Level) {
	std.Level = level
}

// SetOutput sets output stream
func SetOutput(out io.Writer) {
	std.SetOutput(out)
}

// FileWriter creates the path if not exist
func FileWriter(filename string) (*os.File, error) {
	dir := path.Dir(filename)
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(dir, 0755)
		} else {
			return nil, err
		}
	}

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// Println same with Info
func Println(v ...interface{}) {
	if len(loggers) == 0 {
		std.Info(v...)
	} else {
		for _, logger := range loggers {
			logger.Info(v...)
		}
	}
}

// Printf same with Infof
func Printf(format string, v ...interface{}) {
	if len(loggers) == 0 {
		std.Infof(format, v...)
	} else {
		for _, logger := range loggers {
			logger.Infof(format, v...)
		}
	}
}

func Debug(v ...interface{}) {
	if len(loggers) == 0 {
		std.Debug(v...)
	} else {
		for _, logger := range loggers {
			logger.Debug(v...)
		}
	}

}

func Debugf(format string, v ...interface{}) {
	if len(loggers) == 0 {
		std.Debugf(format, v...)
	} else {
		for _, logger := range loggers {
			logger.Debugf(format, v...)
		}
	}

}

func Info(v ...interface{}) {
	if len(loggers) == 0 {
		std.Info(v...)
	} else {
		for _, logger := range loggers {
			logger.Info(v...)
		}
	}
}

func Infof(format string, v ...interface{}) {
	if len(loggers) == 0 {
		std.Infof(format, v...)
	} else {
		for _, logger := range loggers {
			logger.Infof(format, v...)
		}
	}
}

func Warn(v ...interface{}) {
	if len(loggers) == 0 {
		std.Warn(v...)
	} else {
		for _, logger := range loggers {
			logger.Warn(v...)
		}
	}
}

func Warnf(format string, v ...interface{}) {
	if len(loggers) == 0 {
		std.Warnf(format, v...)
	} else {
		for _, logger := range loggers {
			logger.Warnf(format, v...)
		}
	}
}

func Error(v ...interface{}) {
	if len(loggers) == 0 {
		std.Error(v...)
	} else {
		for _, logger := range loggers {
			logger.Error(v...)
		}
	}
}

func Errorf(format string, v ...interface{}) {
	if len(loggers) == 0 {
		std.Errorf(format, v...)
	} else {
		for _, logger := range loggers {
			logger.Errorf(format, v...)
		}
	}
}

func Panic(v ...interface{}) {
	if len(loggers) == 0 {
		std.Panic(v...)
	} else {
		s := ""
		for _, logger := range loggers {
			s = logger.PanicNoPanic(v...)
		}
		panic(s)
	}
}

func Panicf(format string, v ...interface{}) {
	if len(loggers) == 0 {
		std.Panicf(format, v...)
	} else {
		s := ""
		for _, logger := range loggers {
			s = logger.PanicfNoPanic(format, v...)
		}
		panic(s)
	}
}

func Fatal(v ...interface{}) {
	if len(loggers) == 0 {
		std.Fatal(v...)
	} else {
		for _, logger := range loggers {
			logger.FatalNoExit(v...)
		}
		os.Exit(1)
	}
}

func Fatalf(format string, v ...interface{}) {
	if len(loggers) == 0 {
		std.Fatalf(format, v...)
	} else {
		for _, logger := range loggers {
			logger.FatalfNoExit(format, v...)
		}
		os.Exit(1)
	}
}
