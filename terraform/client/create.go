package main

import (
	"context"
	pb "github.com/johess40/AlphagRpc/terraform/proto"
	"log"
)

func createTerra(c pb.TerraformServiceClient) string {
	log.Printf("Create Terra was invoked")

	terra := pb.TerraRequest{
		Provider:     "azurerm",
		ResourceName: "resource_group",
		NumberNeeded: "3",
		Prefix:       "app",
		Tags: map[string]string{
			"env":      "Test",
			"owner":    "DevOps",
			"location": "westus2",
		},
	}

	res, err := c.CreateTerra(context.Background(), &terra)
	if err != nil {
		log.Fatal(err)
	}

	return res.Id
}
