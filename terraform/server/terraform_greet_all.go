package main

import (
	"fmt"
	pb "github.com/johess40/AlphagRpc/terraform/proto"
	"io"
	"log"
	"strconv"
)

func (s *Server) TerraStreamAll(stream pb.TerraformService_TerraStreamAllServer) error {
	log.Printf("Streaming all terraform: %v", stream)

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatal(err)
		}

		atoi, err := strconv.Atoi(req.NumberNeeded)
		if err != nil {
			return err
		}

		for i := 0; i < atoi; i++ {
			res := &pb.TerraResponse{
				Provider:            req.Provider,
				ResourceName:        req.ResourceName,
				NumberToBeRequested: req.NumberNeeded,
				PlannedNamespace:    fmt.Sprintf("%v-%v-%v", req.Prefix, req.ResourceName, strconv.Itoa(i)),
				TagsToBeApplied:     req.Tags,
			}

			err = stream.Send(res)
			if err != nil {
				log.Fatal(err)
			}
		}

	}

}
