package main

import (
	"context"
	pb "github.com/johess40/AlphagRpc/terraform/proto"
	"log"
	"time"
)

func StreamClientTerra(tr pb.TerraformServiceClient) {
	log.Printf("Starting StreamClient For Terraform Service")

	req := []*pb.TerraRequest{
		{
			Provider:     "azurerm",
			ResourceName: "resource_group",
			NumberNeeded: "3",
			Prefix:       "app",
			Tags: map[string]string{
				"env":      "Test",
				"owner":    "DevOps",
				"location": "westus2",
			},
		},
		{
			Provider:     "azurerm",
			ResourceName: "resource_group",
			NumberNeeded: "2",
			Prefix:       "api",
			Tags: map[string]string{
				"env":      "Test",
				"owner":    "DevOps",
				"location": "westus2",
			},
		},
		{
			Provider:     "azurerm",
			ResourceName: "resource_group",
			NumberNeeded: "5",
			Prefix:       "backend",
			Tags: map[string]string{
				"env":      "Test",
				"owner":    "DevOps",
				"location": "westus2",
			},
		},
		{
			Provider:     "azurerm",
			ResourceName: "resource_group",
			NumberNeeded: "1",
			Prefix:       "backend",
			Tags: map[string]string{
				"env":      "Test",
				"owner":    "DevOps",
				"location": "westus2",
			},
		},
	}

	stream, err := tr.TerraLongStream(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range req {
		err := stream.Send(v)

		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(5 * time.Second)

	}
	res, errRec := stream.CloseAndRecv()
	if errRec != nil {
		log.Fatal(err)
	}

	log.Printf("%v", res)
}
