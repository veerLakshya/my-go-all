package mylogger

import (
	"bytes"
	"errors"
	"fmt"
)

type Level int8

var errUnmarshalNilLevel = errors.New("can't unmarshal a nil *Level")

const (
	DebugLevel Level = iota - 1 //-1
	InfoLevel                   //0
	WarnLevel
	ErrorLevel
	DPanicLevel
	PanicLevel
	FatalLevel

	_minLevel = DebugLevel
	_maxLevel = FatalLevel

	InvalidLevel = _maxLevel + 1
)

func ParseLevel(text string) (Level, error) {
	var level Level
	err := level.UnmarshalText([]byte(text))
	return level, err
}

func (l *Level) UnmarshalText(text []byte) error {
	if l == nil {
		return errUnmarshalNilLevel
	}
	if !l.unmarshalText(text) && !l.unmarshalText(bytes.ToLower(text)) {
		return fmt.Errorf("unrecognized level: %q", text)
	}
	return nil
}

func (l *Level) unmarshalText(text []byte) bool {
	switch string(text) {
	case "debug":
		*l = DebugLevel
	case "info":
		*l = InfoLevel
	case "warn", "warning":
		*l = WarnLevel
	case "error":
		*l = ErrorLevel
	case "dpanic":
		*l = DPanicLevel
	case "panic":
		*l = PanicLevel
	case "fatal":
		*l = FatalLevel
	default:
		return false
	}
	return true
}
