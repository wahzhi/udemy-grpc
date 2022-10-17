package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/wahzhi/grpc-go-course/greet/proto"

)

func doFactor(c pb.FactorServiceClient) {   // GreetServiceClient is prototype in pb to handle client request
	                                      // it is an object representing  a client request
										  
	fmt.Println("input number:")
	var  innumber int32
	_, err := fmt.Scanf("%d", &innumber)

	if err != nil {
		log.Fatal(err)
	}
	request := &pb.FactorRequest{
		Number: innumber,
	}

	log.Printf("Client: do Factor was invoked with request.number = %d\n", request.Number)


	str, err := c.Factor(context.Background(), request) //invoke server API 

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes %v\n", err)
	}
   
	for {

		msg, err := str.Recv()     // ********* the most important function to recieve stream package.

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream %v\n", err)
		}

		log.Printf("Response from server to client: %d\n", msg.Factor)  //Result defined in server API



	}

}