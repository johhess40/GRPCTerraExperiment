package main

import (
	"context"
	"fmt"
	pb "github.com/johess40/AlphagRpc/greet/proto"
	"log"
	"strconv"
)

func DoGreet(c pb.GreetServiceClient) {
	response, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName:  "John Hession",
		LastName:   "Hession",
		Location:   "Seattle",
		Occupation: "DevOps",
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Printf(fmt.Sprintf("%v", response))
}

func GetSum(c pb.SumServiceClient) {
	response, err := c.Sum(context.Background(), &pb.SumRequest{
		Sum: 5 + 5,
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Sum is: %v", response.Result)
}

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