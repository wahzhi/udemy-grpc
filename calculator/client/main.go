package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/wahzhi/udemy-grpc/calculator/proto"
    "log"
)
	

var addr string = "localhost:50051"

func main(){
	// responsible to make a connection to the API server defined in /server/main.go
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Client: Failed to connect: %v\n", err)

	}
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)
	//doGreet(c)
	//doFactor(c)
	doAvg(c)

}

