package main

import (
	"flag"
	"fmt"
	"memmi/config"
	"memmi/factory"
	"net/http"
)

func main() {
	cDir := flag.String("cDir", "./config", "Config search directory.")
	cFile := flag.String("cFile", "dev", "Default configuration file name (excluding file extension).")
	flag.Parse()
	config.LoadFromFile(*cDir, *cFile)
	fact := factory.HardCodedFactory{}
	router := fact.GetRouter()
	http.HandleFunc("/", router.GetHandleFunc())
	fmt.Printf("Server listening on '%s'\n", config.AppConfig().Server)
	http.ListenAndServe(config.AppConfig().Server, nil)
}
