package client

import (
	"context"
	"log"
	"net"

	"github.com/353solutions/building/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	addr := ":9292"

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("error: can't listen - %s", err)
	}

	srv := grpc.NewServer()
	var u Buildings
	pb.RegisterBuildingServiceServer(srv, &u)
	reflection.Register(srv)

	log.Printf("info: server ready on %s", addr)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("error: can't serve - %s", err)
	}
}

func (r *Buildings) Start(ctx context.Context, req *pb.StartRequest) (*pb.StartResponse, error) {
	// TODO: Validate req
	resp := pb.StartResponse{
		Id: req.Id,
	}

	// TODO: Work (insert to database ...)
	return &resp, nil
}

type Buildings struct {
	pb.UnimplementedBuildingServiceServer
}
