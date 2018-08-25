package endpoints

import (
	"github.com/sdaf47/go-knowledge-base/small_programms/task/service"
	"github.com/go-kit/kit/endpoint"
	"context"
)

func MakeGetTasksHandler(s service.TaskManagerService) endpoint.Endpoint {
	type responseTask struct {
		Name    string `json:"name"`
		Created string `json:"created"`
		Id      int    `json:"id"`
		Status  string `json:"status"`
	}

	var responseTasks []responseTask

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		tasks, err := s.Tasks()
		if err != nil {
			return
		}

		responseTasks = make([]responseTask, 0, 100)
		for id, task := range tasks {
			responseTasks = append(responseTasks, responseTask{
				Name:    task.Name,
				Status:  task.Status,
				Id:      id,
				Created: task.Created.Format("2006-01-02 15:04:05"),
			})
		}

		response = responseTasks

		return
	}
}
