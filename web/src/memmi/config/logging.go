package config

import (
	"github.com/spf13/viper"
)

func genLogging() Logging {
	logConfig := Logging{}
	logConfig.Format = viper.GetString("logging.format")
	logConfig.SamplingRate = viper.GetFloat64("logging.samplingRate")
	logConfig.Endpoint = viper.GetString("logging.endpoint")
	logConfig.Level = viper.GetString("logging.level")
	switch viper.GetString("logging.type") {
	case "console":
		logConfig.Type = LoggingType_console
	default:
		logConfig.Type = LoggingType_console
	}
	return logConfig
}
