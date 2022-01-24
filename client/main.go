package main

import (
    "context"
    "net/http"
	"fmt"

   // pb "github.com/twitchtv/twirp-example/rpc/helloworld"
   "github.com/jbabinec1/hello-world-service/rpc"
)



// Run the implementation in a local server
func main() {
	fmt.Println("Test Client")

	client := rpc.NewAccountsServiceProtobufClient("http://localhost:6666", &http.Client{})
	
	AccountID, err := client.IsAccountLinked(context.Background(), &rpc.AccountLookupRequest{id: "jared", "pubg"})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Account ID: ", AccountID.String())

	ItemDrop, err := client.PushToAccount(context.Background(), &rpc.AccountDropsResponse{})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Item granted: ", dread.String())
}