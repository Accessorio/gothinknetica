package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "go-core-4/homework14/messenger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewMessengerClient(conn)
	client.Send(context.Background(), &pb.Message{Id: 1, Text: "Hello", Date: "2023-08-30"})
	client.Send(context.Background(), &pb.Message{Id: 2, Text: "world", Date: "2023-08-30"})
	client.Send(context.Background(), &pb.Message{Id: 3, Text: "!", Date: "2023-08-30"})
	err = getAllMessages(context.Background(), client)
	if err != nil {
		fmt.Println("error=", err)
	}
}

func getAllMessages(ctx context.Context, client pb.MessengerClient) error {
	stream, err := client.Messages(context.Background(), &pb.Empty{})
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			book, err := stream.Recv()
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}

			fmt.Printf("Message: %v\n", book)
		}
	}
}
