package config

import (
	"github.com/spf13/viper"
)

func SetDefaults() {
	viper.SetDefault("app.server", ":8081")

	viper.SetDefault("router.logType", "console")
	viper.SetDefault("router.logEndpoint", "localhost:9090")
	viper.SetDefault("router.handlers", []string{"cardSet", "card"})

	viper.SetDefault("cardManagement.type", "inmemory")
	viper.SetDefault("cardManagement.endpoint", "localhost:9091")
	viper.SetDefault("cardManagement.seed", false)

	viper.SetDefault("userManagement.type", "inmemory")
	viper.SetDefault("userManagement.seed", false)

	viper.SetDefault("factory.useSingletons", true)
}
