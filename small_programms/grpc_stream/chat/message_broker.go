package chat

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/sdaf47/go-knowledge-base/small_programms/grpc_stream/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"time"
)

type messageBroker struct {
	sessionRepo SessionRepo

	clients  map[string]*brokerClient
	messages chan stream.Message
	ctx      context.Context
	cancel   context.CancelFunc
}

func NewMessageBroker(ctx context.Context, sessionRepo SessionRepo) *messageBroker {
	broker := &messageBroker{
		clients:     make(map[string]*brokerClient),
		messages:    make(chan stream.Message, 10),
		sessionRepo: sessionRepo,
	}
	broker.ctx, broker.cancel = context.WithCancel(ctx)

	go broker.listenAndServe()

	return broker
}

func (s *messageBroker) Close() {
	s.cancel()
}

func (s *messageBroker) listenAndServe() {
	for {
		select {
		case msg := <-s.messages:
			// broadcast
			for _, client := range s.clients {
				client.replies <- msg
			}
		case <-s.ctx.Done():
			return
		}
	}
}

func (s *messageBroker) Subscribe(ctx context.Context, logon *stream.Logon) (*stream.LogonStatus, error) {
	token, err := s.sessionRepo.Authorize(logon.Username, logon.Password)
	if err != nil {
		return &stream.LogonStatus{
			Error:   err.Error(),
			Success: false,
		}, nil
	}

	p, _ := peer.FromContext(ctx)
	client := newClient(s.ctx, token, logon.Username, p.Addr.String())

	go s.handleClient(client)

	return &stream.LogonStatus{
		Success: true,
		Token:   token,
	}, nil
}

func (s *messageBroker) handleClient(client *brokerClient) {
	s.clients[client.token] = client

	for {
		select {
		case request := <-client.requests:
			ts, err := ptypes.TimestampProto(time.Now().UTC())
			if err != nil {
				panic(err)
			}

			s.messages <- stream.Message{
				Message:   request.Message,
				Username:  client.username,
				Timestamp: ts,
			}
			logrus.
				WithField("username", client.username).
				WithField("message", request.Message).
				Info("new message")
		case <-client.ctx.Done():
			return
		}
	}
}

func (s *messageBroker) OpenStream(conn stream.MessageBroker_OpenStreamServer) error {
	md, ok := metadata.FromIncomingContext(conn.Context())
	if !ok {
		return fmt.Errorf("expect headers")
	}

	var token string
	if len(md.Get("authorization")) <= 0 {
		return fmt.Errorf("expect token in header")
	}
	token = md.Get("authorization")[0]
	client := s.clients[token]
	defer delete(s.clients, token)

	go func() {
		for {
			request, err := conn.Recv()
			if err != nil {
				return
			}

			client.requests <- *request
		}
	}()

	for {
		select {
		case reply := <-client.replies:
			err := conn.Send(&reply)
			if err != nil {
				return err
			}
		}
	}
}
