package main

import (
	"fmt"
	"net"

	"github.com/ashilesh/grpc-stream/chat"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Server Initiated")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic("error while listening")
	}

	grpcServer := grpc.NewServer()

	s := chat.Server{}

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		panic("unable to start grpc server")
	}

}
