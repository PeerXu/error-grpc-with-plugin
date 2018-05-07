package main

import (
	"context"
	"net"

	"google.golang.org/grpc"

	pb "github.com/PeerXu/error-grpc-with-plugin/proto"
)

type greetService struct{}

func (srv *greetService) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	return &pb.GreetResponse{"hello, world"}, nil
}

func Serve() {
	lis, _ := net.Listen("tcp", "0.0.0.0:13401")
	srv := &greetService{}
	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, srv)
	s.Serve(lis)
}
