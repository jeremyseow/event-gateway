package utils

import (
	"github.com/spf13/viper"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type AllUtilityClients struct {
	Logger *zap.Logger
}

func NewAllUtilityClients(conf *viper.Viper) *AllUtilityClients {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	logger, err := loggerConfig.Build(zap.Fields(zap.String("service", viper.GetString("serviceName")), zap.String("serviceVersion", viper.GetString("serviceVersion"))))
	if err != nil {
		panic(err)
	}

	return &AllUtilityClients{
		Logger: logger,
	}
}
