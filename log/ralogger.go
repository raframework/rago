package log

import (
	"log"
)

// Log levels in RFC5424
const (
	LevelEmergency     = 0
	LevelAlert         = 1
	LevelCritical      = 2
	LevelError         = 3
	LevelWarning       = 4
	LevelNotice        = 5
	LevelInformational = 6
	LevelDebug         = 7

	DefaultLevel = LevelInformational
)

var levelNameMap = map[int]string{
	LevelEmergency:     "Emergency",
	LevelAlert:         "Alert",
	LevelCritical:      "Critical",
	LevelError:         "Error",
	LevelWarning:       "Warning",
	LevelNotice:        "Notice",
	LevelInformational: "Informational",
	LevelDebug:         "Debug",
}

type RaLogger struct {
	level int
}

func NewRaLogger() *RaLogger {
	return &RaLogger{
		level: DefaultLevel,
	}
}

func (l *RaLogger) WithLevel(level int) {
	l.level = level
}

func (l *RaLogger) Emergency(message interface{}) {
	l.addRecord(LevelEmergency, message)
}

func (l *RaLogger) Alert(message interface{}) {
	l.addRecord(LevelAlert, message)
}

func (l *RaLogger) Critical(message interface{}) {
	l.addRecord(LevelAlert, message)
}

func (l *RaLogger) Error(message interface{}) {
	l.addRecord(LevelError, message)
}

func (l *RaLogger) Warning(message interface{}) {
	l.addRecord(LevelWarning, message)
}

func (l *RaLogger) Notice(message interface{}) {
	l.addRecord(LevelNotice, message)
}

func (l *RaLogger) Informational(message interface{}) {
	l.addRecord(LevelInformational, message)
}

func (l *RaLogger) Debug(message interface{}) {
	l.addRecord(LevelDebug, message)
}

func (l *RaLogger) addRecord(level int, message interface{}) {
	if level <= l.level {
		log.Println("["+levelNameMap[level]+"]", message)
	}

}
