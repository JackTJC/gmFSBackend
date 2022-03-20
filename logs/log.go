package logs

import (
	"time"

	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

func InitLog() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	sugar = logger.Sugar()
}

func sync() {
	for range time.Tick(time.Minute) {
		sugar.Sync()
	}
}

func Infof(format string, args ...interface{}) {
	sugar.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	sugar.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	sugar.Errorf(format, args...)
}

func FatalF(format string, args ...interface{}) {
	sugar.Errorf(format, args...)
}
