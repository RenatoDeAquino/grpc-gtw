package services

import (
	"context"
	pb "grpc-sample/protos"
	"log"
)

type server struct{}

func NewServer() *server {
	return &server{}
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println("request: ", in.Name)
	return &pb.HelloReply{Message: "hello, " + in.Name}, nil
}
