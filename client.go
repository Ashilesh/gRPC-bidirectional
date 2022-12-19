package main

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/ashilesh/grpc-stream/chat"
	"google.golang.org/grpc"
)

func main() {
	var wg sync.WaitGroup

	fmt.Println("client started")

	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		panic("cannot make connection with localhost:9000")
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

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

	wg.Wait()

}
