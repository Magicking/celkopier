package main

import (
	"log"
	"net"

	pb "github.com/Magicking/celkopier/kopier"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement celkopier.server.
type server struct{}

// Scan implements celkopier.Scan
func (s *server) Scan(ctx context.Context, in *empty.Empty) (*empty.Empty, error) {
	log.Printf("TODO Do the scan\n")
	return in, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterScannerServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
