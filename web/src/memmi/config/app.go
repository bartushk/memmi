package config

import (
	"github.com/spf13/viper"
)

func genApp() App {
	newApp := App{}
	newApp.Server = viper.GetString("app.server")
	return newApp
}
