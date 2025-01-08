// Package logger provides utilities for centralized logging using the zap library.
package logger

// Imports necessary libraries for logging configuration.
import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// log is the global logger instance.
	log *zap.Logger
	
	// LOG_OUTPUT defines the output destination for logs, e.g., stdout or a file path.
	LOG_OUTPUT = "LOG_OUTPUT"
	
	// LOG_LEVEL defines the logging level, e.g., info, error, debug.
	LOG_LEVEL  = "LOG_LEVEL"
)

// init initializes the zap logger with configurations for encoding, output, and log level.
func init() {
	logConfig := zap.Config{
		OutputPaths: []string{getOutputLogs()},
		Level:       zap.NewAtomicLevelAt(getLevelLogs()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	log, _ = logConfig.Build()
}

// Info logs a message with the info level.
// message: The message to be logged.
// tags: Additional context fields for the log.
func Info(message string, tags ...zap.Field) {
	log.Info(message, tags...)
	log.Sync()
}

// Error logs a message with the error level.
// message: The message to be logged.
// err: The error object to be included in the log.
// tags: Additional context fields for the log.
func Error(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Info(message, tags...)
	log.Sync()
}

// getOutputLogs retrieves the log output destination from the environment variable LOG_OUTPUT.
// Returns the destination as a string (default: stdout).
func getOutputLogs() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv(LOG_OUTPUT)))
	if output == "" {
		return "stdout"
	}
	return output
}

// getLevelLogs retrieves the log level from the environment variable LOG_LEVEL.
// Returns the corresponding zapcore.Level (default: InfoLevel).
func getLevelLogs() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(LOG_LEVEL))) {
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	case "debug":
		return zapcore.DebugLevel
	default:
		return zapcore.InfoLevel
	}
}
