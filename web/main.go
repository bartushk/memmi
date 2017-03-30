package main

import (
	"flag"
	"github.com/op/go-logging"
	"memmi/config"
	"memmi/core"
	"memmi/factory"
	"net/http"
)

var log = logging.MustGetLogger("memmi")

func main() {
	cDir := flag.String("cDir", "./config", "Config search directory.")
	cFile := flag.String("cFile", "dev", "Default configuration file name (excluding file extension).")
	flag.Parse()
	config.LoadFromFile(*cDir, *cFile)
	core.InitLogging()
	fact := factory.HardCodedFactory{}
	router := fact.GetRouter()
	http.HandleFunc("/", router.GetHandleFunc())
	log.Infof("Server listening with: \n'%s'\n", config.GetConfig())
	http.ListenAndServe(config.GetConfig().Server, nil)
}
