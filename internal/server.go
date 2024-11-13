package internal

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/ashupednekar/simple-otlp-server/internal/otlp"
	"google.golang.org/grpc"
)

type Server struct {
  pb.UnimplementedTraceServiceServer
}

func (s *Server) Export (
  ctx context.Context, 
  request *pb.ExportTraceServiceRequest,
) (*pb.ExportTraceServiceResponse, error){
  fmt.Printf("%v\n", request)
  return &pb.ExportTraceServiceResponse{}, nil
}

func StartServer(){
  ln, err := net.Listen("tcp", ":4317")
  if err != nil {
    log.Fatalf("error starting collector: %v", err)
  }

  s := grpc.NewServer()
  pb.RegisterTraceServiceServer(s, &Server{})

  log.Printf("gRPC server listening at %v", ln.Addr())

  if err := s.Serve(ln); err != nil {
    log.Fatalf("failed to start gRPC server: %v", err)
  }

}
