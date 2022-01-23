package paperclipserver

import (
    "context"

    pb "github.com/jbabinec1/hello-world-test-service/rpc"
)

// Server implements the UniversalPaperclips service
type Server struct {
    
	
}

// NewServer creates an instance of our server
func NewServer() *Server {
    return &Server{
        
    }
}

// GetPaperclips returns the current paperclip count
func (s *Server) IsAccountLinked(ctx context.Context, accountReq *pb.AccountLookupRequest) (*pb.AccountLookupResponse, error) {
    return &pb.AccountLookupResponse{
        items: "cool item",
    }, nil
}


// CalculateUniverseLifespan calulcates the universe's lifespan and returns some marginally relaxing value
/*  func (s *Server) CalculateUniverseLifespan(ctx context.Context, empty *pb.Empty) (*pb.Dread, error) {
    return &pb.Dread{
        Paperclips:       s.PaperClips,
        UniverseLifespan: "42",
    }, nil */



}