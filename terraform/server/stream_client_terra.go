package main

import (
	pb "github.com/johess40/AlphagRpc/terraform/proto"
	"io"
	"log"
)

func (s *Server) TerraLongStream(stream pb.TerraformService_TerraLongStreamServer) error {
	res := ""
	for {
		req, err := stream.Recv()

		log.Printf("Received data: %v\n", req)

		res += "Resource ID:" + req.Provider + req.ResourceName

		if err == io.EOF {
			return stream.SendAndClose(&pb.TerraResponse{
				Provider:            "",
				ResourceName:        "",
				NumberToBeRequested: "",
				PlannedNamespace:    "",
				TagsToBeApplied:     nil,
			})
		}

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Received data: %v\n", req)
	}
}
