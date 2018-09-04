package transport

import (
	"github.com/sdaf47/go-knowledge-base/small_programms/linkshort/service"
	"github.com/sdaf47/go-knowledge-base/small_programms/linkshort/pb"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/sdaf47/go-knowledge-base/small_programms/linkshort/endpoints"
	"context"
	oldcontext "golang.org/x/net/context"
	"github.com/go-kit/kit/endpoint"
	"google.golang.org/grpc"
	"errors"
)

type grpcServer struct {
	decode grpctransport.Handler
	encode grpctransport.Handler
}

func NewGRPCServer(s service.ShortLinkService) pb.ShortLinkServiceServer {
	return &grpcServer{
		decode: grpctransport.NewServer(
			endpoints.MakeDecodeLinkHandler(s),
			func(_ context.Context, grpcRequest interface{}) (interface{}, error) {
				req := grpcRequest.(*pb.DecodeRequest)

				return endpoints.DecodeRequest{
					Code: req.Code,
				}, nil
			},
			func(_ context.Context, response interface{}) (interface{}, error) {
				resp := response.(endpoints.DecodeResponse)

				return &pb.DecodeReply{Link: resp.Link, Error: resp.Error}, nil
			},
		),
		encode: grpctransport.NewServer(
			endpoints.MakeEncodeLinkHandler(s),
			func(_ context.Context, grpcRequest interface{}) (interface{}, error) {
				req := grpcRequest.(*pb.EncodeRequest)

				return endpoints.EncodeRequest{
					Link: req.Link,
				}, nil
			},
			func(_ context.Context, response interface{}) (interface{}, error) {
				resp := response.(endpoints.EncodeResponse)

				return &pb.EncodeReply{Code: resp.Code, Error: resp.Error}, nil
			},
		),
	}
}

func (s *grpcServer) Decode(ctx oldcontext.Context, req *pb.DecodeRequest) (*pb.DecodeReply, error) {
	_, rep, err := s.decode.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DecodeReply), nil
}

func (s *grpcServer) Encode(ctx oldcontext.Context, req *pb.EncodeRequest) (*pb.EncodeReply, error) {
	_, rep, err := s.encode.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.EncodeReply), nil
}

type grpcClient struct {
	encode endpoint.Endpoint
	decode endpoint.Endpoint
}

func NewGRPCClient(conn *grpc.ClientConn) service.ShortLinkService {
	var decodeEndpoint endpoint.Endpoint
	{
		decodeEndpoint = grpctransport.NewClient(
			conn,
			"pb.ShortLinkService",
			"Decode",
			func(_ context.Context, request interface{}) (interface{}, error) {
				req := request.(endpoints.DecodeRequest)
				return &pb.DecodeRequest{Code: req.Code}, nil
			},
			func(_ context.Context, grpcReply interface{}) (interface{}, error) {
				reply := grpcReply.(*pb.DecodeReply)
				return endpoints.DecodeResponse{Link: reply.Link}, nil
			},
			pb.DecodeReply{},
		).Endpoint()
	}

	var encodeEndpoint endpoint.Endpoint
	{
		encodeEndpoint = grpctransport.NewClient(
			conn,
			"pb.ShortLinkService",
			"Encode",
			func(_ context.Context, request interface{}) (interface{}, error) {
				req := request.(endpoints.EncodeRequest)
				return &pb.EncodeRequest{Link: req.Link}, nil
			},
			func(_ context.Context, grpcReply interface{}) (interface{}, error) {
				reply := grpcReply.(*pb.EncodeReply)
				return endpoints.EncodeResponse{Code: reply.Code, Error: reply.Error}, nil
			},
			pb.EncodeReply{},
		).Endpoint()
	}

	return &grpcClient{
		decode: decodeEndpoint,
		encode: encodeEndpoint,
	}
}

func (c *grpcClient) Encode(link string) (code string, err error) {
	resp, err := c.encode(context.Background(), endpoints.EncodeRequest{
		Link: link,
	})
	if err != nil {
		return
	}

	response := resp.(endpoints.EncodeResponse)
	return response.Code, makeErr(response.Error)
}

func (c *grpcClient) Decode(code string) (link string, err error) {
	resp, err := c.decode(context.Background(), endpoints.DecodeRequest{
		Code: code,
	})
	if err != nil {
		return
	}

	response := resp.(endpoints.DecodeResponse)
	return response.Link, makeErr(response.Error)
}

func makeErr(mess string) error {
	if mess == "" {
		return nil
	}
	return errors.New(mess)
}
