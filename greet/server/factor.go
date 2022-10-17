package main

import (
	
	"log"

	pb "github.com/wahzhi/grpc-go-course/greet/proto"
)

// this implements server side end point, the GreetManyTimes interface required by the grpc
//     func for the object Server,  in = the request message
//         here the reuest messge has 3 fields.
// this exact function is called from the client c.greet(...) endpoint
//  The grpc handles the "middleware"




func (s *Server)  Factor(in *pb.FactorRequest, stream pb.FactorService_FactorServer) error{
	log.Printf("Server: Greet many times function was invoked with %v\n", in )

	var   k int32 = 2
	var  N int32 

	N = in.Number
	
	for N>1 {
		
		if (N % k) == 0  {
	//		res := fmt.Sprintf("factor is %d\n", k)
			stream.Send(&pb.FactorResponse{
				Factor: k,
			}) 
			N = N / k 
			 // if k evenly divides into N
			   // divide N by k so that we have the rest of the number left.
		} else
			{k = k + 1}
	}
	
	return nil;
}
	
	