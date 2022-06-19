package main

import (
	"context"
	pb "github.com/johess40/AlphagRpc/terraform/proto"
	"log"
)

func (s *Server) Terra(ctx context.Context, in *pb.TerraRequest) (*pb.TerraResponse, error) {
	log.Printf("Sum was received %v\n", in)
	return &pb.TerraResponse{
		Provider:            in.Provider,
		ResourceName:        in.ResourceName,
		NumberToBeRequested: in.NumberNeeded,
		PlannedNamespace:    in.Prefix,
		TagsToBeApplied:     in.Tags,
	}, nil
}
