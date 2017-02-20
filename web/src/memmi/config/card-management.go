package config

import (
	"github.com/spf13/viper"
)

func genCardManagement() CardManagement {
	cMan := CardManagement{}
	cMan.Endpoint = viper.GetString("cardManagement.endpoint")
	switch viper.GetString("cardManagement.type") {
	case "inmemory":
		cMan.Type = ManagementType_inmemory
	default:
		cMan.Type = ManagementType_inmemory
	}
	return cMan
}
