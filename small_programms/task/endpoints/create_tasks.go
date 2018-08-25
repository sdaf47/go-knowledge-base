package endpoints

import (
	"github.com/sdaf47/go-knowledge-base/small_programms/task/service"
	"github.com/go-kit/kit/endpoint"
	"context"
)

type CreateTaskRequest struct {
	Name string
}

func MakeCreateTaskHandler(s service.TaskManagerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		_, err = s.CreateTask(request.(CreateTaskRequest).Name)
		if err != nil {
			return
		}

		return
	}
}
