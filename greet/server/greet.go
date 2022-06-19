package main

import (
	"context"
	"fmt"
	pb "github.com/johess40/AlphagRpc/greet/proto"
	"log"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet was received %v\n", in)
	return &pb.GreetResponse{
		ResultFirstName:  fmt.Sprintf("First Name: %v\n", in.LastName),
		ResultLastName:   fmt.Sprintf("Last Name: %v\n", in.LastName),
		ResultLocation:   fmt.Sprintf("Located At: %v\n", in.Location),
		ResultOccupation: fmt.Sprintf("Occupation is: %v\n", in.Occupation),
	}, nil
}

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum was received %v\n", in)
	return &pb.SumResponse{
		Result: in.Sum,
	}, nil
}

func (s *Server) Terra(ctx context.Context, in *pb.TerraRequest) (*pb.TerraResponse, error) {
	return &pb.TerraResponse{
		Provider:            in.Provider,
		ResourceName:        in.ResourceName,
		NumberToBeRequested: in.NumberNeeded,
		PlannedNamespace:    in.Prefix,
		TagsToBeApplied:     in.Tags,
	}, nil
}
