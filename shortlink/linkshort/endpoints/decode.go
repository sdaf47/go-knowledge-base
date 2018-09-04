package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"context"
	"github.com/sdaf47/go-knowledge-base/shortlink/linkshort/service"
)

type DecodeRequest struct {
	Code string `json:"code"`
}

type DecodeResponse struct {
	Link   string `json:"link"`
	Error  string `json:"error"`
}

func MakeDecodeLinkHandler(s service.ShortLinkService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		link, err := s.Decode(request.(DecodeRequest).Code)
		if err != nil {
			return
		}

		response = DecodeResponse{
			Link: link,
		}

		return
	}
}
