package chat

import (
	"context"
	"github.com/sdaf47/go-knowledge-base/small_programms/grpc_stream/pb"
)

type brokerClient struct {
	token    string
	addr     string
	username string
	requests chan stream.Request
	replies  chan stream.Message
	ctx      context.Context
	cancel   context.CancelFunc
}

func newClient(ctx context.Context, token, username, addr string) *brokerClient {
	client := &brokerClient{
		token:    token,
		username: username,
		requests: make(chan stream.Request),
		replies:  make(chan stream.Message),
		addr:     addr,
	}

	client.ctx, client.cancel = context.WithCancel(ctx)

	return client
}
