package logger

import (
	"github.com/doloopwhile/logrusltsv"
	"github.com/sirupsen/logrus"
)

// Formatter is a type for Logger format.
type Formatter struct {
	logrus.Formatter
}

// Formatters is a struct variable for Config.
var Formatters = struct {
	Text Formatter
	JSON Formatter
	LTSV Formatter
}{
	Text: Formatter{Formatter: &logrus.TextFormatter{}},
	JSON: Formatter{Formatter: &logrus.JSONFormatter{}},
	LTSV: Formatter{Formatter: &logrusltsv.Formatter{}},
}
