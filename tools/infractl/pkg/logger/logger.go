package logger

import (
	"io"
	"os"

	"github.com/charmbracelet/log"
)

// LogLevel represents the logging verbosity
type LogLevel int

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)

// Logger provides a structured logging interface with emojis
type Logger struct {
	logger *log.Logger
	level  LogLevel
}

// LoggerInterface defines the contract for logging methods
type LoggerInterface interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	WithFields(fields map[string]any) *Logger
}

// Ensure Logger implements LoggerInterface
var _ LoggerInterface = (*Logger)(nil)

// NewLogger creates a new logger with specified options and emojis
func NewLogger(output io.Writer, level LogLevel) *Logger {
	if output == nil {
		output = os.Stderr
	}

	// Create a new Charmbracelet logger
	charmLogger := log.New(output)

	// Set log level with emoji support
	switch level {
	case LogLevelDebug:
		charmLogger.SetLevel(log.DebugLevel)
	case LogLevelInfo:
		charmLogger.SetLevel(log.InfoLevel)
	case LogLevelWarn:
		charmLogger.SetLevel(log.WarnLevel)
	case LogLevelError:
		charmLogger.SetLevel(log.ErrorLevel)
	default:
		charmLogger.SetLevel(log.InfoLevel)
	}

	// Configure styling
	charmLogger.SetFormatter(log.TextFormatter)
	charmLogger.SetReportTimestamp(false)

	return &Logger{
		logger: charmLogger,
		level:  level,
	}
}

// DefaultLogger creates a logger with default settings and emojis
func DefaultLogger() *Logger {
	return NewLogger(os.Stderr, LogLevelInfo)
}

// WithFields adds structured fields to the logger
func (l *Logger) WithFields(fields map[string]any) *Logger {
	newLogger := l.logger.With(fields)
	return &Logger{
		logger: newLogger,
		level:  l.level,
	}
}

// Debug logs a debug message
func (l *Logger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

// Info logs an info message
func (l *Logger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

// Warn logs a warning message
func (l *Logger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

// Error logs an error message
func (l *Logger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}
