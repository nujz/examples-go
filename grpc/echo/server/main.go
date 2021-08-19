package main

import (
	"context"
	"net"

	pb "github.com/nujz/examples-go/grpc/echo/protobuf"
	"google.golang.org/grpc"
)

type EchoServer struct {
	pb.UnimplementedEchoServiceServer
}

func (s *EchoServer) Test(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{Msg: in.Msg}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	gs := grpc.NewServer()
	pb.RegisterEchoServiceServer(gs, &EchoServer{})

	if err := gs.Serve(lis); err != nil {
		panic(err)
	}
}
