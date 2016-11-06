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
	s := fmt.Sprint(v...)
	std.Output(s)
}

func Alert(v ...interface{}) {
	s := fmt.Sprint(v...)
	std.Output(s)
}

func Critical(v ...interface{}) {
	s := fmt.Sprint(v...)
	std.Output(s)
}

func Error(v ...interface{}) {
	s := fmt.Sprint(v...)
	std.Output(s)
}

func Warning(v ...interface{}) {
	s := fmt.Sprint(v...)
	std.Output(s)
}

func Notice(v ...interface{}) {
	s := fmt.Sprint(v...)
	std.Output(s)
}

func Informational(v ...interface{}) {
	s := fmt.Sprint(v...)
	std.Output(s)
}

func Debug(v ...interface{}) {
	s := fmt.Sprint(v...)
	std.Output(s)
}

func Emergencyf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	std.Output(s)
}

func Alertf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	std.Output(s)
}

func Criticalf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	std.Output(s)
}

func Errorf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	std.Output(s)
}

func Warningf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	std.Output(s)
}

func Noticef(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	std.Output(s)
}

func Informationalf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	std.Output(s)
}

func Debugf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	std.Output(s)
}
