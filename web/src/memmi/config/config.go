package config

import (
	"github.com/spf13/viper"
)

var cMan CardManagement
var uMan UserManagement
var router Router
var app App

func LoadFromFile(directory, filename string) {
	SetDefaults()
	viper.SetConfigName(filename)
	viper.AddConfigPath(directory)
	viper.ReadInConfig()

	app = genApp()
	cMan = genCardManagement()
	uMan = genUserManagement()
	router = genRouter()
}

func Load() {
	LoadFromFile("~/", ".memmirc")
}

func CardManagementConfig() CardManagement {
	return cMan
}

func UserManagementConfig() UserManagement {
	return uMan
}

func RouterConfig() Router {
	return router
}

func AppConfig() App {
	return app
}