package core

import (
	"github.com/op/go-logging"
	"memmi/config"
	"os"
)

func InitLogging() {
	logConfig := config.LoggingConfig()
	format := logging.MustStringFormatter(logConfig.Format)
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatted := logging.NewBackendFormatter(backend, format)
	backendLeveled := logging.AddModuleLevel(backendFormatted)
	switch logConfig.Level {
	case "info":
		backendLeveled.SetLevel(logging.INFO, "")
		break
	case "notice":
		backendLeveled.SetLevel(logging.NOTICE, "")
		break
	case "warning":
		backendLeveled.SetLevel(logging.WARNING, "")
		break
	case "err":
		backendLeveled.SetLevel(logging.ERROR, "")
		break
	case "crit":
		backendLeveled.SetLevel(logging.CRITICAL, "")
		break
	}

	logging.SetBackend(backendLeveled)
}
