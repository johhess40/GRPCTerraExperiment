package main

import (
	"context"
	pb "github.com/johess40/AlphagRpc/terraform/proto"
	"log"
)

func readTerra(t pb.TerraformServiceClient, id string) *pb.TerraResponse {
	log.Printf("-----Reading Terraform Blob----")

	req := &pb.TerraId{
		Id: id,
	}

	res, err := t.ReadTerra(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Found: %v\n", res)

	return res
}
