package main

import (
	"context"
	pb "github.com/johess40/AlphagRpc/terraform/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"net"
)

var addr = "0.0.0.0:50051"

var collection *mongo.Collection

type Server struct {
	pb.TerraformServiceServer
}

type TerraBlob struct {
	ID           primitive.ObjectID `bson:"_id, omitempty"`
	Provider     string             `bson:"provider"`
	ResourceName string             `bson:"resource_name"`
	Prefix       string             `bson:"prefix"`
	NumberNeeded string             `bson:"number_needed"`
	Tags         map[string]string  `bson:"tags"`
}

func main() {
	RegisterServices(addr)
}

func DocumentTerra(terra *pb.TerraRequest) *pb.TerraResponse {
	return &pb.TerraResponse{
		Id:                  terra.Id,
		Provider:            terra.Provider,
		ResourceName:        terra.ResourceName,
		NumberToBeRequested: terra.NumberNeeded,
		PlannedNamespace:    terra.ResourceName,
		TagsToBeApplied:     terra.Tags,
	}
}

func RegisterServices(st string) {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("terra_state_db").Collection("alpha_terra")
	lis, err := net.Listen("tcp", st)
	if err != nil {
		log.Fatalf("Failed to listen on address %v\n", addr)
	}

	log.Printf("Listen on address %v\n", addr)
	s := grpc.NewServer()
	//reflection.Register(s)

	pb.RegisterTerraformServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Error %v\n", err)
	}
}
