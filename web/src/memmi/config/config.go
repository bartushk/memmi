package config

import (
	"github.com/spf13/viper"
)

var cMan CardManagement
var uMan UserManagement
var log Logging
var router Router
var app App
var fact Factory

func LoadFromFile(directory, filename string) {
	SetDefaults()
	viper.SetConfigName(filename)
	viper.AddConfigPath(directory)
	viper.ReadInConfig()

	cMan = genCardManagement()
	uMan = genUserManagement()
	router = genRouter()
	fact = genFactory()
	log = genLogging()
	app = genApp()
	app.CardMan = &cMan
	app.UserMan = &uMan
	app.Router = &router
	app.Factory = &fact
	app.Logging = &log
}

func Load() {
	LoadFromFile("~/", ".memmirc")
}

func GetConfig() App {
	return app
}
