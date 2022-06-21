package main

import (
	"context"
	pb "github.com/johess40/AlphagRpc/terraform/proto"
	"io"
	"log"
	"strconv"
)

func DisplayTerra(t pb.TerraformServiceClient) {
	terra, err := t.Terra(context.Background(), &pb.TerraRequest{
		Provider:     "azurerm",
		ResourceName: "resource_group",
		NumberNeeded: "3",
		Prefix:       "app",
		Tags: map[string]string{
			"env":      "Test",
			"owner":    "DevOps",
			"location": "westus2",
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	GeneratePlan(terra)

}

func GeneratePlan(t *pb.TerraResponse) {
	n, err := strconv.Atoi(t.NumberToBeRequested)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Provider %v will be registered...\n", t.Provider)

	for i := 0; i < n; i++ {
		log.Printf("Resource %v_%v with name %v-%v-%v will be created...\n", t.Provider, t.ResourceName, t.ResourceName, strconv.Itoa(i), t.PlannedNamespace)
	}

	for k, v := range t.TagsToBeApplied {
		log.Printf("Tag with Key: %v, Value: %v will be applied...\n", k, v)
	}
}

func StreamTerra(t pb.TerraformServiceClient) {
	terra, err := t.TerraMany(context.Background(), &pb.TerraRequest{
		Provider:     "azurerm",
		ResourceName: "resource_group",
		NumberNeeded: "3",
		Prefix:       "app",
		Tags: map[string]string{
			"env":      "Test",
			"owner":    "DevOps",
			"location": "westus2",
		},
	})

	if err != nil {
		log.Fatal(err)
	}
	for {
		msg, err := terra.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		n, err := strconv.Atoi(msg.NumberToBeRequested)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Provider %v will be registered...\n", msg.Provider)

		for i := 0; i < n; i++ {
			log.Printf("Resource %v_%v with name %v-%v-%v will be created...\n", msg.Provider, msg.ResourceName, msg.ResourceName, strconv.Itoa(i), msg.PlannedNamespace)
		}

		for k, v := range msg.TagsToBeApplied {
			log.Printf("Tag with Key: %v, Value: %v will be applied...\n", k, v)
		}

	}
}
