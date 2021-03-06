package main

import (
	"context"
	"log"
	"time"

	"github.com/talento90/grpc-calculator/client/calculator"
	"google.golang.org/grpc"
)

func main() {
	const address = "127.0.0.1:5000"

	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error connecting to grpc server: %v", err)
	}

	defer conn.Close()

	c := calculator.NewCalculatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Division(ctx, &calculator.OperationRequest{
		FirstNumber:  100,
		SecondNumber: 20,
	})

	if err != nil {
		log.Fatalf("Error : %v", err)
	}

	log.Printf("Result: %f", r.Result)
}