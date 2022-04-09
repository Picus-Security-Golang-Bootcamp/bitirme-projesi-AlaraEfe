package logger

import (
	"fmt"
	"github.com/AlaraEfe/BasketService/packages/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//NewLogger function creates a new logger with given log level
func NewLogger(config *config.Config) {
	logLevel, err := zapcore.ParseLevel(config.Logger.Level)
	if err != nil {
		panic(fmt.Sprintf("Unknown log level %v", logLevel))
	}

	var cfg zap.Config
	if config.Logger.Development {
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		cfg = zap.NewProductionConfig()
	}

	logger, err := cfg.Build()
	if err != nil {
		logger = zap.NewNop()
	}
	zap.ReplaceGlobals(logger)

}

//Calls global logger and sync it
func Close() {
	defer zap.L().Sync()
}
