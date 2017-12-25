package log

import (
	"os"
	"testing"
)

// Use default global logger
func TestOneLogger(t *testing.T) {
	// Default output to stderr
	SetLevel(LevelDebug)
	Debug("stderr: Debug")
	Debugf("stderr: Debug%s", "f")
	Info("stderr: Info")
	Infof("stderr: Info%s", "f")
	Warn("stderr: Warn")
	Warnf("stderr: Warn%s", "f")
	Error("stderr: Error")
	Errorf("stderr: Error%s", "f")
	//Panic("stderr: Panic")
	//Panicf("stderr: Panic%s", "f")
	//Fatal("stderr: Fatal")
	//Fatalf("stderr: Fatal%s", "f")

	// Change output to file
	fw, err := FileWriter("logs/test.log")
	if err != nil {
		panic(err)
		return
	}
	SetOutput(fw)
	SetLevel(LevelDebug)
	Info("file: Info")
}

// New self logger instance
func TestMultiLoggers(t *testing.T) {
	// New logger for output stderr
	logStd := New("", LevelDebug, os.Stderr)
	logStd.Info("a stderr logger")

	// New logger for output file
	fw, err := FileWriter("logs/test.log")
	if err != nil {
		panic(err)
		return
	}
	logFile := New("", LevelDebug, fw)
	logFile.Info("a file logger")
}

// Add multi handlers to default global
func TestMultiHandlers(t *testing.T) {
	// New logger for output stderr
	logStd := New("", LevelDebug, os.Stderr)
	AddHandler(logStd)

	// New logger for output file
	fw, err := FileWriter("logs/test.log")
	if err != nil {
		panic(err)
	}
	logFile := New("", LevelDebug, fw)
	AddHandler(logFile)

	// Test
	Debug("MultiHandlers: Debug")
	Info("MultiHandlers: Info")
	Warn("MultiHandlers: Warn")
	Error("MultiHandlers: Error")
	//Panic("MultiHandlers: Panic")
	//Fatal("MultiHandlers: Fatal")
}

// Test Rotate wirter
// By day: RotateTimeByDay
// By hour: RotateTimeByHour
func Test(t *testing.T) {
	rw, err := NewRotateWriter("./logs/data.log", RotateTimeClose)
	if err != nil {
		panic(err)
	}
	SetOutput(rw)
	Info("TEST INFO")
}
