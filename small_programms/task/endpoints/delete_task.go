package endpoints

import (
	"github.com/sdaf47/go-knowledge-base/small_programms/task/service"
	"github.com/go-kit/kit/endpoint"
	"context"
	"strconv"
)

func MakeDeleteTaskHandler(s service.TaskManagerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		vars := ctx.Value("vars").(map[string]string)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			return
		}

		err = s.DeleteTask(id)

		return
	}
}
