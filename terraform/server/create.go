package main

import (
	"context"
	pb "github.com/johess40/AlphagRpc/terraform/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) CreateTerra(c context.Context, in *pb.TerraRequest) (*pb.TerraId, error) {
	log.Printf("Creating Terraform Object: %v\n", in)

	data := pb.TerraRequest{
		Provider:     in.Provider,
		ResourceName: in.ResourceName,
		NumberNeeded: in.NumberNeeded,
		Prefix:       in.Prefix,
		Tags:         in.Tags,
	}

	_, err := collection.InsertOne(c, &data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			err.Error(),
		)
	}

	return &pb.TerraId{Id: in.Id}, nil
}
