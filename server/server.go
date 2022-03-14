package main

import (
	pb "calculator/gen/proto"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedAddServiceServer
}

func (s *server) Add(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	a, b := req.GetA(), req.GetB()
	result := a + b
	return &pb.Response{Result: result}, nil
}

func (s *server) Mul(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	a, b := req.GetA(), req.GetB()
	result := a * b
	return &pb.Response{Result: result}, nil

}

func main() {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)

	}

	srv := grpc.NewServer()

	pb.RegisterAddServiceServer(srv, &server{})

	//for serealizing and deserealization
	reflection.Register(srv)

	if e := srv.Serve(lis); e != nil {
		panic(err)
	}

}
