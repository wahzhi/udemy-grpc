package main

import (
	"context"
	
	"log"

	pb "github.com/wahzhi/udemy-grpc/calculator/proto"
)
	

// doAvg takes the interface c as input
func doAvg(c pb.CalculatorServiceClient) {
   log.Println("doAvg was invoked")

   stream, err := c.Avg(context.Background())

   if err != nil {
	log.Fatalf("Error while openning the stream %v\n", err)
   }

   numbers := []int32{3,5,9,54,23}

   for _, number := range numbers {
	log.Printf("Sending numbers: %d\n", number)
	stream.Send(&pb.AvgRequest{
		Number: number,
	})
   }
   res, err := stream.CloseAndRecv()   // close the stream and ready to recv from server

   if err != nil {
	log.Fatalf("Error while receiving response: %v\n", err)
   }

   log.Printf("Avg: %f\n", res.Result)

}
