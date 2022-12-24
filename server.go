package main

import (
	"fmt"
	"net"
	"time"

	"github.com/ashilesh/grpc-stream/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func main() {
	fmt.Println("Server Initiated")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic("error while listening")
	}

	enforcement := keepalive.EnforcementPolicy{
		MinTime:             10 * time.Second,
		PermitWithoutStream: true,
	}

	grpcServer := grpc.NewServer(grpc.KeepaliveEnforcementPolicy(enforcement),
		grpc.KeepaliveParams(keepalive.ServerParameters{

			// MaxConnectionAgeGrace: 30 * time.Second,
		}))

	s := chat.Server{}

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		panic("unable to start grpc server")
	}

}
