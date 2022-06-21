package main

import (
	"context"
	pb "github.com/johess40/AlphagRpc/terraform/proto"
	"io"
	"log"
	"time"
)

func StreamAllTerra(p pb.TerraformServiceClient) {
	log.Printf("Streaming all Terraform requests...")

	stream, err := p.TerraStreamAll(context.Background())

	if err != nil {
		log.Fatal(err)
	}

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

	waitterra := make(chan struct{})

	go func() {
		for _, v := range req {
			log.Printf("Sending Terraform request: %v:", v)
			err = stream.Send(v)
			if err != nil {
				return
			}
			time.Sleep(1 * time.Second)

		}
		err = stream.CloseSend()
		if err != nil {
			return
		}
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatal(err)
			}

			log.Printf("%v", res)
		}
		close(waitterra)
	}()

	<-waitterra
}
