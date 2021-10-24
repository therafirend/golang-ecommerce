package zapLog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)


var log *zap.Logger

func init() {
	var err error

	zc := zap.NewProductionConfig()
	zec := zap.NewProductionEncoderConfig()
	zec.TimeKey = "time"
	zec.EncodeTime = zapcore.ISO8601TimeEncoder
	zec.StacktraceKey = ""
	zc.EncoderConfig = zec

	log, err = zc.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(msg string,fields ...zap.Field) {
	log.Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field){
	log.Debug(msg, fields...)
}

func Error(msg string, fields ...zap.Field){
	log.Error(msg, fields...)
}
