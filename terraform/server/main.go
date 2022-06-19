package main

import (
	pb "github.com/johess40/AlphagRpc/terraform/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var addr = "0.0.0.0:50051"

type Server struct {
	pb.TerraformServiceServer
}

func main() {
	RegisterServices(addr)
}

func RegisterServices(st string) {
	lis, err := net.Listen("tcp", st)
	if err != nil {
		log.Fatalf("Failed to listen on address %v\n", addr)
	}

	log.Printf("Listen on address %v\n", addr)
	s := grpc.NewServer()

	pb.RegisterTerraformServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Error %v\n", err)
	}
}
