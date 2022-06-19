package main

import (
	"context"
	"fmt"
	pb "github.com/johess40/AlphagRpc/greet/proto"
	"log"
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
