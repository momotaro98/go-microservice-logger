package logger

import (
	"github.com/sirupsen/logrus"
)

const (
	// XTransactionID is a key of transaction ID for microservice.
	// This is supposed to be set in HTTP header and SQS attributes and so on.
	// This variable is used in each service.
	XTransactionID = "X-Transaction-ID"

	logKeyOfXTxID       = "request-id"
	logKeyOfServiceName = "service-name"
)

var (
	// ServiceName is supposed to be embedded by using
	// ldflags at build time of a service application.
	ServiceName = "not-set"

	// defaultLog is a private instance of this package.
	//
	// The default setting as Microservice rule
	// Formatter: LTSV
	// Out:       stdout
	// MinLevel:  Info
	defaultLog = defaultLogger()
)

// Logger is public interface of this package.
type Logger interface {
	Debug(xTxID interface{}, msg string, fields ...Field)
	Info(xTxID interface{}, msg string, fields ...Field)
	Warn(xTxID interface{}, msg string, fields ...Field)
	Error(xTxID interface{}, msg string, fields ...Field)
	Panic(xTxID interface{}, msg string, fields ...Field)
}

func defaultLogger() Logger {
	return NewLogger(NewConfig())
}

// AlterDefaultLogger changes default logger by the given config.
// This affects logger.Debug, Info, Warn, Error methods.
func AlterDefaultLogger(conf *Config) {
	defaultLog = NewLogger(conf)
}

// Debug logs a message at level Debug as microservice default logger
func Debug(xTxID interface{}, msg string, fields ...Field) {
	defaultLog.Debug(xTxID, msg, fields...)
}

// Info logs a message at level Info as microservice default logger
func Info(xTxID interface{}, msg string, fields ...Field) {
	defaultLog.Info(xTxID, msg, fields...)
}

// Warn logs a message at level Warn as microservice default logger
func Warn(xTxID interface{}, msg string, fields ...Field) {
	defaultLog.Warn(xTxID, msg, fields...)
}

// Error logs a message at level Error as microservice default logger
func Error(xTxID interface{}, msg string, fields ...Field) {
	defaultLog.Error(xTxID, msg, fields...)
}

// Panic logs a message at level Panic as microservice default logger
func Panic(xTxID interface{}, msg string, fields ...Field) {
	defaultLog.Panic(xTxID, msg, fields...)
}

type logger struct {
	*logrus.Logger
	config *Config
}

// NewLogger is a constructor of Logger interface.
// Basically, the rule is that microservice applications should use
// Log global instances without using this method.
// You should call this method with setting up Config when you do
// development or research.
func NewLogger(config *Config) Logger {
	var l = logrus.New()
	l.Level, _ = logrus.ParseLevel(config.minLevel.String())
	l.Formatter = config.formatter
	l.Out = config.out

	return &logger{
		Logger: l,
		config: config,
	}
}

// Debug logs a message at level Debug.
func (l *logger) Debug(xTxID interface{}, msg string, fields ...Field) {
	if l.enabledLogLevel(Levels.Debug) {
		l.withFields(l.mergeToFields(xTxID, fields...)...).Debug(msg)
	}
}

// Info logs a message at level Info.
func (l *logger) Info(xTxID interface{}, msg string, fields ...Field) {
	if l.enabledLogLevel(Levels.Info) {
		l.withFields(l.mergeToFields(xTxID, fields...)...).Info(msg)
	}
}

// Warn logs a message at level Warn.
func (l *logger) Warn(xTxID interface{}, msg string, fields ...Field) {
	if l.enabledLogLevel(Levels.Warn) {
		l.withFields(l.mergeToFields(xTxID, fields...)...).Warn(msg)
	}
}

// Error logs a message at level Error.
func (l *logger) Error(xTxID interface{}, msg string, fields ...Field) {
	if l.enabledLogLevel(Levels.Error) {
		l.withFields(l.mergeToFields(xTxID, fields...)...).Error(msg)
	}
}

// Panic logs a message at level Panic.
func (l *logger) Panic(xTxID interface{}, msg string, fields ...Field) {
	if l.enabledLogLevel(Levels.Panic) {
		l.withFields(l.mergeToFields(xTxID, fields...)...).Panic(msg)
	}
}

func (l *logger) mergeToFields(xTxID interface{}, fields ...Field) []Field {
	// Required fields
	var requiredFields = []Field{
		{key: logKeyOfXTxID, value: xTxID},
		{key: logKeyOfServiceName, value: ServiceName},
	}

	mFields := make([]Field, 0, len(requiredFields)+len(fields))
	mFields = append(mFields, requiredFields...)
	return append(mFields, fields...)
}

func (l *logger) withFields(fields ...Field) *logrus.Entry {
	var cnvFields = make(logrus.Fields, len(fields))
	for _, f := range fields {
		cnvFields[f.key] = f.value
	}
	return l.Logger.WithFields(cnvFields)
}

func (l *logger) enabledLogLevel(level Level) bool {
	return l.config.minLevel <= level
}
