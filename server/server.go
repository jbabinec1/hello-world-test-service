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

// GetPaperclips returns the current paperclip count
func (s *Server) IsAccountLinked(ctx context.Context, accountReq *pb.AccountLookupRequest) (*pb.AccountLookupResponse, error) {
    return &pb.AccountLookupResponse{
        items: "cool item",
    }, nil
}


}