package main

import (
	"memmi/handlers"
	"net/http"
)

func main() {
	router := new(handlers.Router)
	http.HandleFunc("/", router.GetHandleFunc())
	http.ListenAndServe(":8080", nil)
}
