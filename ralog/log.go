package ralog

import (
	"fmt"
	"os"
)

var std = NewLogger(os.Stderr)

func SetLevel(level LogLevel) {
	std.WithLevel(level)
}

// Emergency log message on these conditions.
// System is unusable.
func Emergency(v ...interface{}) {
	std.Emergency(v...)
}

// Alert log message on these conditions.
// Action must be taken immediately.
//
// Example: Entire website down, database unavailable, etc. This should
// trigger the SMS alerts and wake you up.
func Alert(v ...interface{}) {
	std.Alert(v...)
}

// Critical log message on these conditions.
// Critical conditions.
//
// Example: Application component unavailable, unexpected exception.
func Critical(v ...interface{}) {
	std.Critical(v...)
}

// Error log message on these conditions.
// Runtime errors that do not require immediate action but should typically
// be logged and monitored.
func Error(v ...interface{}) {
	std.Error(v...)
}

// Warning log message on these conditions.
// Exceptional occurrences that are not errors.
//
// Example: Use of deprecated APIs, poor use of an API, undesirable things
// that are not necessarily wrong.
func Warning(v ...interface{}) {
	std.Warning(v...)
}

// Notice log message on these conditions.
// Normal but significant events.
func Notice(v ...interface{}) {
	std.Notice(v...)
}

// Informational log message on these conditions.
// Interesting events.
//
// Example: User logs in, SQL logs.
func Informational(v ...interface{}) {
	std.Informational(v...)
}

// Debug log message on these conditions.
// Detailed debug information.
func Debug(v ...interface{}) {
	std.Debug(v...)
}

func Emergencyf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	std.Emergency(s)
}

func Alertf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	std.Alert(s)
}

func Criticalf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	std.Critical(s)
}

func Errorf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	std.Error(s)
}

func Warningf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	std.Warning(s)
}

func Noticef(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	std.Notice(s)
}

func Informationalf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	std.Informational(s)
}

func Debugf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	std.Debug(s)
}
