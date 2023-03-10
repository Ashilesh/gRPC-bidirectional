package chat

import (
	"fmt"
	"io"
	// "time"
)

type Server struct{}

func (s *Server) Connect(stream ChatService_ConnectServer) error {
	fmt.Println("connect called")
	go func() {
		// time.Sleep(9 * time.Second)
		stream.Send(&Message{ChatMessage: "from server msg 1"})
		// time.Sleep(10 * time.Second)
		stream.Send(&Message{ChatMessage: "from server msg 2"})
		fmt.Println("data sent")
	}()

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("no more data")
			return err
		}

		if err != nil {
			fmt.Println("error while receiving data", err)
			return err
		}

		fmt.Println("Server : ", req.ChatMessage)
	}
}

func (s *Server) mustEmbedUnimplementedChatServiceServer() {}
