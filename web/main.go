package main

import (
	"fmt"
	"memmi/handlers"
	"net/http"
)

func main() {
	server := ":8080"
	router := new(handlers.Router)
	router.Logger = new(handlers.ConsoleLogger)
	router.Authenticator = new(handlers.MockAuthenticator)
	http.HandleFunc("/", router.GetHandleFunc())
	fmt.Printf("Server listening on '%s'\n", server)
	http.ListenAndServe(":8080", nil)
}
