package main

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/ashilesh/grpc-stream/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func main() {
	var wg sync.WaitGroup

	kc := keepalive.ClientParameters{
		Time:                10 * time.Second,
		Timeout:             2 * time.Second,
		PermitWithoutStream: true,
	}

	fmt.Println("client started")

	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure(), grpc.WithKeepaliveParams(kc))
	if err != nil {
		panic("cannot make connection with localhost:9000")
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)
	ctx := context.Background()
	// defer cancel()

	res, err := c.Connect(ctx)
	if err != nil {
		panic("error while calling RPC")
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			m, err := res.Recv()
			if err == io.EOF {
				fmt.Println("end EOF")
				return
			}
			if err != nil {
				panic(err)
			}

			fmt.Println("client", m)
		}
	}()

	res.Send(&chat.Message{ChatMessage: "this is initial message"})
	res.Send(&chat.Message{ChatMessage: "this is 2 message"})
	fmt.Println("data sent")

	// time.Sleep(10 * time.Second)
	// res.Send(&chat.Message{ChatMessage: "this is delayed message"})
	// time.Sleep(10 * time.Second)
	// res.Send(&chat.Message{ChatMessage: "this is delayed message"})
	wg.Wait()
}
