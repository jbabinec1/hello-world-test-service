package main

import (
	"fmt"
	"log"
	"net/http"

 
	"github.com/jbabinec1/hello-world-test-service/server"
	"github.com/jbabinec1/hello-world-test-service/rpc"

)

func main() {
	fmt.Printf("Starting Service on :6666")

	server := server.NewServer()
	twirpHandler := rpc.NewAccountsServiceServer(server, nil)
    fmt.Println(twirpHandler)
	log.Fatal(http.ListenAndServe(":6666", twirpHandler))
}
