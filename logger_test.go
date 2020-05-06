package logger

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	const (
		xTranVal = "request-1234567890"
	)
	var (
		buff = new(bytes.Buffer)
		cnf  = NewConfig(
			WithOut(buff),
			WithMinLevel(Levels.Debug),
			// Default formatter is LTSV
		)
		log = NewLogger(cnf)
	)

	tests := []struct {
		name     string
		logFunc  func()
		expected string
	}{
		{
			name: "Debug Log",
			logFunc: func() {
				log.Debug(xTranVal,
					"debug message",
					F("id", 123),
					E(errors.New("err text")))
			},
			expected: fmt.Sprintf("level:debug\tmsg:debug message\terror:err text\tid:123\t%s:%s\t%s:%s",
				logKeyOfXTxID, xTranVal,
				logKeyOfServiceName, ServiceName),
		},
		{
			name: "Info Log",
			logFunc: func() {
				log.Info(xTranVal,
					"info message",
					F("id", 123),
					E(errors.New("err text")))
			},
			expected: fmt.Sprintf("level:info\tmsg:info message\terror:err text\tid:123\t%s:%s\t%s:%s",
				logKeyOfXTxID, xTranVal,
				logKeyOfServiceName, ServiceName),
		},
		{
			name: "Warn Log",
			logFunc: func() {
				log.Warn(xTranVal,
					"warn message",
					F("id", 123),
					E(errors.New("err text")))
			},
			expected: fmt.Sprintf("level:warning\tmsg:warn message\terror:err text\tid:123\t%s:%s\t%s:%s",
				logKeyOfXTxID, xTranVal,
				logKeyOfServiceName, ServiceName),
		},
		{
			name: "Error Log",
			logFunc: func() {
				log.Error(xTranVal,
					"error message",
					F("id", 123),
					E(errors.New("err text")))
			},
			expected: fmt.Sprintf("level:error\tmsg:error message\terror:err text\tid:123\t%s:%s\t%s:%s",
				logKeyOfXTxID, xTranVal,
				logKeyOfServiceName, ServiceName),
		},

		// [Note] We don't do test for Panic method because logrus.Panic executes panic() in it.
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buff.Reset()
			// Act
			test.logFunc()
			// Assert
			assert.Equal(t, test.expected+"\n", buff.String()[31:])
			// 31 means removing time:RFC3339 field like "time:2020-04-08T19:42:00+09:00"
			// This needs TZ="Asia/Tokyo" on machine because of "+09:00" chars.
			// [Note] LTSV formatter makes sure that time:RFC3339 is first field of the defaultLog output.
		})
	}
}
