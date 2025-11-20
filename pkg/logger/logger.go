package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var globalLogger *zap.Logger

// Config holds logger configuration
type Config struct {
	Level      string // debug, info, warn, error
	Format     string // json, console
	OutputPath string
}

// InitLogger initializes the global logger
func InitLogger(config Config) (*zap.Logger, error) {
	level := parseLogLevel(config.Level)

	var zapConfig zap.Config
	if config.Format == "json" || config.Format == "" {
		zapConfig = zap.NewProductionConfig()
	} else {
		zapConfig = zap.NewDevelopmentConfig()
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	zapConfig.Level = zap.NewAtomicLevelAt(level)

	if config.OutputPath != "" {
		zapConfig.OutputPaths = []string{config.OutputPath}
	} else {
		zapConfig.OutputPaths = []string{"stdout"}
	}

	zapConfig.ErrorOutputPaths = []string{"stderr"}

	logger, err := zapConfig.Build(
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	if err != nil {
		return nil, err
	}

	globalLogger = logger
	return logger, nil
}

// GetLogger returns the global logger instance
func GetLogger() *zap.Logger {
	if globalLogger == nil {
		// Create a default logger if not initialized
		logger, _ := zap.NewProduction()
		globalLogger = logger
	}
	return globalLogger
}

// WithFields creates a child logger with additional fields
func WithFields(fields map[string]interface{}) *zap.Logger {
	logger := GetLogger()
	zapFields := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		zapFields = append(zapFields, zap.Any(k, v))
	}
	return logger.With(zapFields...)
}

// WithRequestID creates a child logger with request ID
func WithRequestID(requestID string) *zap.Logger {
	return GetLogger().With(zap.String("request_id", requestID))
}

// With creates a child logger with additional zap fields
func With(fields ...zap.Field) *zap.Logger {
	return GetLogger().With(fields...)
}

// Sync flushes any buffered log entries
func Sync() error {
	if globalLogger != nil {
		return globalLogger.Sync()
	}
	return nil
}

func parseLogLevel(levelStr string) zapcore.Level {
	switch strings.ToLower(levelStr) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn", "warning":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

func init() {
	// Initialize with a default production logger
	if globalLogger == nil {
		logger, err := zap.NewProduction()
		if err != nil {
			// Fallback to a no-op logger
			logger = zap.NewNop()
		}
		globalLogger = logger
	}
}

// Close closes the logger and flushes any buffered entries
func Close() error {
	if globalLogger != nil {
		err := globalLogger.Sync()
		globalLogger = nil
		return err
	}
	return nil
}

// ReplaceGlobals replaces the global logger
func ReplaceGlobals(logger *zap.Logger) {
	zap.ReplaceGlobals(logger)
	globalLogger = logger
}

// Simple convenience methods that use the global logger

// Debug logs a debug message
func Debug(msg string, fields ...zap.Field) {
	GetLogger().Debug(msg, fields...)
}

// Info logs an info message
func Info(msg string, fields ...zap.Field) {
	GetLogger().Info(msg, fields...)
}

// Warn logs a warning message
func Warn(msg string, fields ...zap.Field) {
	GetLogger().Warn(msg, fields...)
}

// Error logs an error message
func Error(msg string, fields ...zap.Field) {
	GetLogger().Error(msg, fields...)
}

// Fatal logs a fatal message and exits
func Fatal(msg string, fields ...zap.Field) {
	GetLogger().Fatal(msg, fields...)
	os.Exit(1)
}
