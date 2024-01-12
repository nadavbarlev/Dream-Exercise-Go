package main

import (
	"fmt"
	"log"
	"net/http"
)

const portNumber string = ":8123"

func main() {

	// Create the server
	server := &http.Server{
		Addr: portNumber,
		Handler: routes(),
	}
	
	// Upload the server
	fmt.Println(fmt.Sprintf("Listening on port %s", portNumber))
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return
	}
}
