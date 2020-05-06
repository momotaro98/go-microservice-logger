package logger

import (
	"fmt"
	"strings"
)

const (
	strDebug   = "debug"
	strError   = "error"
	strInfo    = "info"
	strPanic   = "panic"
	strUnknown = "unknown"
	strWarn    = "warn"
)

// Level is a type of Log level
type Level int8

type levels struct {
	Debug Level
	Info  Level
	Warn  Level
	Error Level
	Panic Level
}

// String returns string of Level type
func (l Level) String() string {
	switch l {
	case Levels.Debug:
		return strDebug
	case Levels.Info:
		return strInfo
	case Levels.Warn:
		return strWarn
	case Levels.Error:
		return strError
	case Levels.Panic:
		return strPanic
	default:
		return strUnknown
	}
}

// Levels provides levels of logging
// Debug Info Warn Error Panic Fatal
var Levels = func() levels {
	const (
		Debug Level = iota + 1
		Info
		Warn
		Error
		Panic
	)
	return levels{
		Debug: Debug,
		Info:  Info,
		Warn:  Warn,
		Error: Error,
		Panic: Panic,
	}
}()

// ParseLevel parses string to Level type
func ParseLevel(level string) Level {
	switch strings.ToLower(level) {
	case strDebug:
		return Levels.Debug
	case strInfo:
		return Levels.Info
	case strWarn:
		return Levels.Warn
	case strError:
		return Levels.Error
	case strPanic:
		return Levels.Panic
	default:
		panic(fmt.Sprintf("unknown defaultLog level: %s", level))
	}
}
