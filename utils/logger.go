package utils

import (
	"log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerUtils struct{}

var LoggerUtil = LoggerUtils{}

var Log *zap.SugaredLogger

func (l *LoggerUtils) Init() {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	logger, err := loggerConfig.Build()
	if err != nil {
		log.Fatal(err)
	}

	Log = logger.Sugar()

	defer Log.Sync()
}
