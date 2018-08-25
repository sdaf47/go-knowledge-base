package transport

import (
	kithttp "github.com/go-kit/kit/transport/http"
	"net/http"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"reflect"
	"github.com/sdaf47/go-knowledge-base/small_programms/task/endpoints"
	"github.com/sdaf47/go-knowledge-base/small_programms/task/service"
)

func MakeHandler(s service.TaskManagerService) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
		kithttp.ServerBefore(getVarsMiddleware),
	}

	var createTaskHandler http.Handler
	{
		createTaskHandler = kithttp.NewServer(
			endpoints.MakeCreateTaskHandler(s),
			newJSONRequestDecoder(func() interface{} { return &endpoints.CreateTaskRequest{} }),
			encodeResponse,
			opts...
		)
		createTaskHandler = recoveringMiddleware(createTaskHandler)
	}

	var getTasksHandler http.Handler
	{
		getTasksHandler = kithttp.NewServer(
			endpoints.MakeGetTasksHandler(s),
			nilRequestDecoder,
			encodeResponse,
			opts...
		)
		getTasksHandler = recoveringMiddleware(getTasksHandler)
	}

	var getTaskHandler http.Handler
	{
		getTaskHandler = kithttp.NewServer(
			endpoints.MakeGetTaskHandler(s),
			nilRequestDecoder,
			encodeResponse,
			opts...
		)
		getTaskHandler = recoveringMiddleware(getTaskHandler)
	}

	var deleteTaskHandler http.Handler
	{
		deleteTaskHandler = kithttp.NewServer(
			endpoints.MakeDeleteTaskHandler(s),
			nilRequestDecoder,
			encodeResponse,
			opts...
		)
		deleteTaskHandler = recoveringMiddleware(deleteTaskHandler)
	}

	var updateTaskHandler http.Handler
	{
		updateTaskHandler = kithttp.NewServer(
			endpoints.MakeUpdateTaskHandler(s),
			newJSONRequestDecoder(func() interface{} { return &endpoints.UpdateTaskRequest{} }),
			encodeResponse,
			opts...
		)
		updateTaskHandler = recoveringMiddleware(updateTaskHandler)
	}

	r := mux.NewRouter()

	// task
	r.Methods("POST").Path("/tasks").Handler(createTaskHandler)
	r.Methods("GET").Path("/tasks").Handler(getTasksHandler)
	r.Methods("GET").Path("/tasks/{id}").Handler(getTaskHandler)
	r.Methods("DELETE").Path("/tasks/{id}").Handler(deleteTaskHandler)
	r.Methods("POST").Path("/tasks/{id}").Handler(updateTaskHandler)

	return r
}

type errorer interface {
	error() error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func recoveringMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				var err error
				switch t := r.(type) {
				case string:
					err = fmt.Errorf(t)
				case error:
					err = t
				}

				encodeError(request.Context(), err, writer)
			}
		}()

		handler.ServeHTTP(writer, request)
	})
}

func getVarsMiddleware(ctx context.Context, r *http.Request) context.Context {
	vars := mux.Vars(r)

	return context.WithValue(ctx, "vars", vars)
}

func newJSONRequestDecoder(requestFactory func() interface{}) kithttp.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		request := requestFactory()
		if err := json.NewDecoder(r.Body).Decode(request); err != nil {
			return nil, err
		}

		refReq := reflect.ValueOf(request)

		return refReq.Elem().Interface(), nil
	}
}

func nilRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	return r, nil
}
