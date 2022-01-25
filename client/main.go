package main

import (
    "context"
    "net/http"
	"fmt"

	"github.com/jbabinec1/hello-world-test-service/rpc"
)



// Run the implementation in a local server
func main() {
	fmt.Println("Test Client")

	client := rpc.NewAccountsServiceProtobufClient("http://localhost:6666", &http.Client{})

	AccountID, err := client.IsAccountLinked(context.Background(), &rpc.AccountLookupRequest{Platform: "console", PlatformID: "xbox"})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Account ID: ", AccountID.String())


	ItemDrop, err := client.PushToAccount(context.Background(), &rpc.AccountDropsRequest{Id:"noob weapon", Game: "pubg"})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Item granted: ", ItemDrop.String())
}