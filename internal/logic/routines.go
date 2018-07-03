package logic

import (
	"log"
	"time"

	pb "github.com/Magicking/celkopier/kopier"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

//const (
//	address = "localhost:50051"
//)

func DoScan(address string) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewScannerClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Scan(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("could not scan: %v", err)
	}
	log.Printf("Do the scan: %x", r)
}
