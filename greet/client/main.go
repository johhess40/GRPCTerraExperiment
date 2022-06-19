package main

import (
	pb "github.com/johess40/AlphagRpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial failed %v\n", err)
	}

	c := pb.NewGreetServiceClient(conn)
	sum := pb.NewSumServiceClient(conn)

	DoGreet(c)
	GetSum(sum)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

}
