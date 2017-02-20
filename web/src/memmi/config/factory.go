package config

import (
	"github.com/spf13/viper"
)

func genFactory() Factory {
	factory := Factory{}
	factory.UseSingletons = viper.GetBool("factory.useSingletons")
	return factory
}
