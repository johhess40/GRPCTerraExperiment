package main

import (
	pb "github.com/johess40/AlphagRpc/terraform/proto"
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

	terra := pb.NewTerraformServiceClient(conn)

	//DisplayTerra(terra)
	//StreamTerra(terra)
	//StreamClientTerra(terra)
	//StreamAllTerra(terra)
	createTerra(terra)
	readTerra(terra, "62b290db8f48ecf3ce0d9414")
	readTerra(terra, "ppohoeiuhrperere")
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

}
