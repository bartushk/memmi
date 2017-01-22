package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type ConsoleLogger struct {
}

func (logger *ConsoleLogger) Log(r *http.Request) {
	fmt.Println("###########################")
	fmt.Printf("Url: %s\n", r.URL.EscapedPath())
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		fmt.Printf("Body: %s\n", body)
	}
	fmt.Println("###########################")
}
