package log

import (
	"fmt"
	"io"
)

// Log levels in RFC5424
const (
	LEmergency     = 0
	LAlert         = 1
	LCritical      = 2
	LError         = 3
	LWarning       = 4
	LNotice        = 5
	LInformational = 6
	LDebug         = 7

	DefaultLevel = LInformational
)

type LogLevel uint8

// String returns a multi-character representation of the log level.
func (l LogLevel) String() string {
	switch l {
	case LEmergency:
		return "Emergency"
	case LAlert:
		return "Alert"
	case LCritical:
		return "Critical"
	case LError:
		return "Error"
	case LWarning:
		return "Warning"
	case LNotice:
		return "Notice"
	case LInformational:
		return "Informational"
	case LDebug:
		return "Debug"
	default:
		panic(fmt.Sprintf("Rago: unhandled loglevel %d", l))
	}
}

type Logger struct {
	level LogLevel  // log level
	out   io.Writer // destination for output
	buf   []byte    // for accumulating text to write
}

func NewLogger(out io.Writer) *Logger {
	return &Logger{
		level: DefaultLevel,
		out:   out,
	}
}

func (l *Logger) WithLevel(level LogLevel) {
	l.level = level
}

func (l *Logger) Emergency(v ...interface{}) {
	l.addRecord(LEmergency, v...)
}

func (l *Logger) Alert(v ...interface{}) {
	l.addRecord(LAlert, v...)
}

func (l *Logger) Critical(v ...interface{}) {
	l.addRecord(LAlert, v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.addRecord(LError, v...)
}

func (l *Logger) Warning(v ...interface{}) {
	l.addRecord(LWarning, v...)
}

func (l *Logger) Notice(v ...interface{}) {
	l.addRecord(LNotice, v...)
}

func (l *Logger) Informational(v ...interface{}) {
	l.addRecord(LInformational, v...)
}

func (l *Logger) Debug(v ...interface{}) {
	l.addRecord(LDebug, v...)
}

func (l *Logger) Output(s string) error {
	l.buf = []byte(s)
	if len(s) == 0 || s[len(s)-1] != '\n' {
		l.buf = append(l.buf, '\n')
	}
	_, err := l.out.Write(l.buf)

	return err
}

func (l *Logger) addRecord(level LogLevel, v ...interface{}) {
	if level <= l.level {
		l.Output("[" + level.String() + "]" + fmt.Sprint(v...))
	}
}
