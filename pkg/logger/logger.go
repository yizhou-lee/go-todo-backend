package logger

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// Log is the global logger for the application.
	Log *zap.Logger
)

func InitLogger(logLevel, logPath string) error {
	// Ensure the log directory exists
	err := os.MkdirAll(filepath.Dir(logPath), 0755)
	if err != nil {
		return err
	}

	// Define log level
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	// Define log configuration
	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(level),
		Development: false,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout", logPath},
		ErrorOutputPaths: []string{"stderr"},
	}

	// Build logger
	Log, err = cfg.Build()
	if err != nil {
		return err
	}

	// Replace global logger
	zap.ReplaceGlobals(Log)

	return nil
}
