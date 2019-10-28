package main

import (
	"github.com/sdaf47/go-knowledge-base/small_programms/grpc_stream/chat"
	"github.com/sdaf47/go-knowledge-base/small_programms/grpc_stream/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":9981")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	defer server.Stop()

	stream.RegisterMessageBrokerServer(server, chat.NewMessageBroker(context.Background(), chat.NewSessionRepo()))

	panic(server.Serve(listener))
}
