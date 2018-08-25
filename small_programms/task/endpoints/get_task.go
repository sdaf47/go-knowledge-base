package endpoints

import (
	"github.com/sdaf47/go-knowledge-base/small_programms/task/service"
	"github.com/go-kit/kit/endpoint"
	"context"
	"strconv"
)

func MakeGetTaskHandler(s service.TaskManagerService) endpoint.Endpoint {
	type responseTask struct {
		Name    string `json:"name"`
		Created string `json:"created"`
		Id      int    `json:"id"`
		Status  string `json:"status"`
	}

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		vars := ctx.Value("vars").(map[string]string)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			return
		}

		task, err := s.Task(id)
		if err != nil {
			return
		}

		response = responseTask{
			Name:    task.Name,
			Status:  task.Status,
			Id:      id,
			Created: task.Created.Format("2006-01-02 15:04:05"),
		}

		return
	}
}
