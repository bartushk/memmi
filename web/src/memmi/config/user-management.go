package config

import (
	"github.com/spf13/viper"
)

func genUserManagement() UserManagement {
	uMan := UserManagement{}
	uMan.Endpoint = viper.GetString("userManagement.endpoint")
	switch viper.GetString("userManagement.type") {
	case "inmemory":
		uMan.Type = ManagementType_inmemory
	default:
		uMan.Type = ManagementType_inmemory
	}
	return uMan
}
