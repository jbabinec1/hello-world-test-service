package server

import (
    "context"

    pb "github.com/jbabinec1/hello-world-test-service/rpc"
)

// Server implements the UniversalPaperclips service
type Server struct {

	AccountLookupResponse string

}

var request

// NewServer creates an instance of our server
func NewServer() *Server {
    return &Server{
       AccountLookupResponse: "jared",
    }
}


func (s *Server) IsAccountLinked(ctx context.Context, request *pb.AccountLookupRequest) (response *pb.AccountLookupResponse, err error) {
    return &pb.AccountLookupResponse{
        AccountID: "jared",
    }, nil
}


func (s *Server) PushToAccount(ctx context.Context, request *pb.AccountDropsRequest) (response *pb.AccountDropsResponse, err error) {
    return &pb.AccountDropsResponse {
		Item: "noob_weapon",
	}, nil
}