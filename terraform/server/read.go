package main

import (
	"context"
	pb "github.com/johess40/AlphagRpc/terraform/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) ReadTerra(ctx context.Context, terra *pb.TerraId) (*pb.TerraResponse, error) {
	log.Printf("Reading Terraform Object...")

	oid, err := primitive.ObjectIDFromHex(terra.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			err.Error())
	}
	data := pb.TerraRequest{}
	filter := bson.M{"_id": oid}

	c := collection.FindOne(ctx, filter)

	if err := c.Decode(&data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			err.Error())
	}

	return DocumentTerra(&data), nil
}
