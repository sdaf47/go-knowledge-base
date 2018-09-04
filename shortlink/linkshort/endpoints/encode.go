package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"context"
	"github.com/sdaf47/go-knowledge-base/small_programms/linkshort/service"
)

type EncodeRequest struct {
	Link string `json:"link"`
}

type EncodeResponse struct {
	Code  string `json:"code"`
	Error string `json:"error"`
}

func MakeEncodeLinkHandler(s service.ShortLinkService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		link, err := s.Encode(request.(EncodeRequest).Link)
		if err != nil {
			return
		}

		response = EncodeResponse{
			Code: link,
		}

		return
	}
}
