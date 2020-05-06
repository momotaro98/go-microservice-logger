package logger

import (
	"github.com/sirupsen/logrus"
)

// Field is a struct type for additional fields of Logging.
type Field struct {
	key   string
	value interface{}
}

// F is a function to make normal Field with key and value.
func F(key string, value interface{}) Field {
	return Field{
		key:   key,
		value: value,
	}
}

// E is a function to make error Field with key and value.
func E(err error) Field {
	return Field{
		key:   logrus.ErrorKey,
		value: err,
	}
}
