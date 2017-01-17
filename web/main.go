package main

import (
	"memmi/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.GetRouter().GetHandleFunc())
	http.ListenAndServe(":8080", nil)
}
