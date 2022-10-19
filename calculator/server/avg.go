package main

import (
	"io"
	"log"

	pb "github.com/wahzhi/udemy-grpc/calculator/proto"
)

	

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {

	log.Println("Avg function was invoked")

	var sum int32 = 0
	count := 0

	for {
		
		req, err := stream.Recv()

		if err == io.EOF {
			avg := float64(sum) / float64(count)
			log.Printf("Server side: averging numbers: %f\n", avg)
			return stream.SendAndClose(&pb.AvgResponse{
				Result: avg,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream  %v\n", err)
		}

		log.Printf("Server side: Receiving  numbers: %d\n", req.Number)
		sum += req.Number
		count++
	}
}