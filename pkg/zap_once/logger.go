package zap_helper

import (
	"sync"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger

	once sync.Once
)

func GetLogger() *zap.Logger {
	once.Do(func() {
		var err error
		logger, err = zap.NewProduction()
		if err != nil {
			panic(err)
		}
	})

	return logger
}
