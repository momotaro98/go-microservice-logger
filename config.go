package logger

import (
	"io"
	"os"
)

// Config is a configuration structure for Logger.
//
// formatter: Format type of logging, TEXT or JSON, or LTSV
// out:       io.Writer of the logger output
// minLevel:  Minimum level to out
type Config struct {
	formatter Formatter
	out       io.Writer
	minLevel  Level
}

// Formatter is a getter method to get formatter of the config.
func (c *Config) Formatter() Formatter {
	return c.formatter
}

// Out is a getter method to get out as io.Writer of the config.
func (c *Config) Out() io.Writer {
	return c.out
}

// MinLevel is a getter method to get minLevel of the config.
func (c *Config) MinLevel() Level {
	return c.minLevel
}

// ConfigOption is a type of function to set the fields of
// Config for Functional Options Pattern.
type ConfigOption func(*Config)

// NewConfig is a constructor of *Config.
//
// Default setting as Microservice rule
// Formatter: LTSV
// Out:       stdout
// MinLevel:  Info
func NewConfig(options ...ConfigOption) *Config {
	config := &Config{
		formatter: Formatters.LTSV,
		out:       os.Stdout,
		minLevel:  Levels.Info,
	}
	for _, option := range options {
		option(config)
	}
	return config
}

// WithMinLevel - specify Logging minimum level.
func WithMinLevel(minLevel Level) ConfigOption {
	return func(c *Config) {
		c.minLevel = minLevel
	}
}

// WithFormatter - specify format type of logger.
func WithFormatter(formatter Formatter) ConfigOption {
	return func(c *Config) {
		c.formatter = formatter
	}
}

// WithOut - specify output destination of logger.
func WithOut(out io.Writer) ConfigOption {
	return func(c *Config) {
		c.out = out
	}
}
