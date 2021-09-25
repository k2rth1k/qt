package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func fields() (m map[string]interface{}) {
	m = make(map[string]interface{})
	m["service_name"] = "quick_trade"
	m["service_ver"] = "v1"
	return m
}
func InitZapLog() *zap.SugaredLogger {
	config := &zap.Config{
		Level:             zap.NewAtomicLevelAt(zapcore.Level(zap.InfoLevel)),
		Encoding:          "json",
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stdout"},
		InitialFields:     fields(),
		DisableStacktrace: false,
		DisableCaller:     false,
		EncoderConfig: zapcore.EncoderConfig{
			CallerKey:      "caller",
			LevelKey:       "level",
			LineEnding:     "\n",
			MessageKey:     "message",
			NameKey:        "logger",
			StacktraceKey:  "stacktrace",
			TimeKey:        "@timestamp",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeName:     zapcore.FullNameEncoder,
			EncodeDuration: zapcore.NanosDurationEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
		},
	}
	logger, _ := config.Build()

	return logger.Sugar()
}
