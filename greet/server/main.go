package main
import (
	"log"
	"net"

	pb "github.com/wahzhi/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct{
	pb.FactorServiceServer  //interface 
}

func main() {
	
	// open the port "addr" and listening to input
	lis, err := net.Listen("tcp",addr)


	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	// instantiate a generic server
	s := grpc.NewServer()

	// register this server to grpc
	pb.RegisterFactorServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

	// the above does nothing... it just establish a connection with the client.
	
	// below is handling any request.
}








