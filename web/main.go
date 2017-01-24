package main

import (
	"fmt"
	"memmi/request"
	"net/http"
)

func main() {
	server := ":8081"
	router := new(request.Router)
	router.Logger = new(request.ConsoleLogger)
	router.Authenticator = new(request.MockAuthenticator)
	http.HandleFunc("/", router.GetHandleFunc())
	fmt.Printf("Server listening on '%s'\n", server)
	http.ListenAndServe(server, nil)
}
