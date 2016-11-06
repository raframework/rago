package log

import (
	"fmt"
	"os"
)

var std = NewLogger(os.Stderr)

func SetLevel(level LogLevel) {
	std.WithLevel(level)
}

func Emergency(v ...interface{}) {
	std.Emergency(v...)
}

func Alert(v ...interface{}) {
	std.Alert(v...)
}

func Critical(v ...interface{}) {
	std.Critical(v...)
}

func Error(v ...interface{}) {
	std.Error(v...)
}

func Warning(v ...interface{}) {
	std.Warning(v...)
}

func Notice(v ...interface{}) {
	std.Notice(v...)
}

func Informational(v ...interface{}) {
	std.Informational(v...)
}

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
