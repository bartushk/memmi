package config

import (
	"github.com/spf13/viper"
)

func SetDefaults() {
	viper.SetDefault("app.server", ":8081")

	viper.SetDefault("router.handlers", []string{"cardSet", "card"})

	viper.SetDefault("cardManagement.type", "inmemory")
	viper.SetDefault("cardManagement.endpoint", "localhost:9091")
	viper.SetDefault("cardManagement.seed", false)

	viper.SetDefault("userManagement.type", "inmemory")
	viper.SetDefault("userManagement.seed", false)

	viper.SetDefault("factory.useSingletons", true)

	viper.SetDefault("logging.type", "console")
	viper.SetDefault("logging.endpoint", "localhost:9090")
	viper.SetDefault("logging.samplingRate", 1.0)
	viper.SetDefault("logging.format",
		"%{color}%{level:.4s} %{time:15:04:05.000} %{shortfunc} %{id:03x}:%{color:reset} %{message}")
}
