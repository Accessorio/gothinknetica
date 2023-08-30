package main

import (
	"context"
	"log"
	"net"
	"sync"

	pb "go-core-4/homework14/messenger"
	"google.golang.org/grpc"
)

type Messages struct {
	Data []pb.Message
	rwm  sync.RWMutex

	pb.UnimplementedMessengerServer
}

func main() {
	msg := &Messages{}
	listener, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMessengerServer(grpcServer, msg)
	grpcServer.Serve(listener)
}

func (m *Messages) Send(_ context.Context, message *pb.Message) (*pb.Empty, error) {
	m.rwm.Lock()
	m.Data = append(m.Data, *message)
	m.rwm.Unlock()
	return new(pb.Empty), nil
}

func (m *Messages) Messages(_ *pb.Empty, stream pb.Messenger_MessagesServer) error {
	m.rwm.RLock()
	for i := range m.Data {
		stream.Send(&m.Data[i])
	}
	m.rwm.RUnlock()
	return nil
}
