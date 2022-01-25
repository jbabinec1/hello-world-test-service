package server

import (
    "context"

    pb "github.com/jbabinec1/hello-world-test-service/rpc"
)

// Server implements the UniversalPaperclips service
type Server struct {

	AccountLookupResponse string

}

// NewServer creates an instance of our server
func NewServer() *Server {
    return &Server{
       AccountLookupResponse: "jared",
    }
}


func (s *Server) IsAccountLinked(ctx context.Context, req *pb.AccountLookupRequest) (*pb.AccountLookupResponse, error) {
    return &pb.AccountLookupResponse{
        AccountID: "jared",
    }, nil
}


func (s *Server) PushToAccount(ctx context.Context, request *pb.AccountDropsRequest) (*pb.AccountDropsResponse, error) {
    return &pb.AccountDropsResponse {
		Item: "noob_weapon",
	}, nil
}