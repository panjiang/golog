package log

import (
	"errors"
	"fmt"
)

// Config struct for configuring logger
type Config struct {
	File   string `json:"file"`
	Level  string `json:"level"`
	Rotate string `json:"rotate"`
}

// DebugString returns a string containing all information
func (c *Config) DebugString() string {
	return fmt.Sprintf("file: %s, level: %s, rotate: %s", c.File, c.Level, c.Rotate)
}

// ParseConfig external method for parsing log config
func ParseConfig(conf *Config) error {
	if conf.File != "" {
		rotateTime, ok := rotateNames[conf.Rotate]
		if !ok {
			return errors.New("invalid rotate name")
		}
		rw, err := NewRotateWriter(conf.File, rotateTime)
		if err != nil {
			return err
		}
		SetOutput(rw)
	}

	// log level
	var level Level = -1
	for k, v := range levelsName {
		if v == conf.Level {
			level = k
			break
		}
	}
	if level == -1 {
		return errors.New("invalid level name")
	}

	SetLevel(level)
	return nil
}
