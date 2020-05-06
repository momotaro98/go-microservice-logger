package logger

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	t.Run("Default Config", func(t *testing.T) {
		// Arrange
		c := NewConfig()
		// Act
		formatter := c.Formatter()
		minLevel := c.MinLevel()
		out := c.Out()
		// Assert
		assert.Equal(t, Formatters.LTSV, formatter)
		assert.Equal(t, Levels.Info, minLevel)
		assert.Equal(t, os.Stdout, out)
	})
}

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name     string
		options  []ConfigOption
		expected *Config
	}{
		{
			name: "Full specified options",
			options: []ConfigOption{
				WithFormatter(Formatters.Text),
				WithOut(os.Stderr),
				WithMinLevel(Levels.Error),
			},
			expected: &Config{
				formatter: Formatters.Text,
				out:       os.Stderr,
				minLevel:  Levels.Error,
			},
		},
		{
			name:    "Empty specified option then return default config",
			options: []ConfigOption{},
			expected: &Config{ // Default setting
				formatter: Formatters.LTSV,
				out:       os.Stdout,
				minLevel:  Levels.Info,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Act
			act := NewConfig(test.options...)
			// Assert
			assert.Equal(t, test.expected, act)
		})
	}
}
