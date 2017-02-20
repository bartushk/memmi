package config

import (
	"github.com/spf13/viper"
)

func genRouter() Router {
	route := Router{}
	route.LogEndpoint = viper.GetString("router.logEndpoint")
	switch viper.GetString("router.logType") {
	case "console":
		route.LogType = LoggingType_console
	default:
		route.LogType = LoggingType_console
	}
	return route
}
