package main

import (
	"github.com/sdaf47/go-knowledge-base/small_programms/task/service"
	"github.com/sdaf47/go-knowledge-base/small_programms/database"
	"github.com/sdaf47/go-knowledge-base/small_programms/task/transport"
	"net/http"
)

func main() {
	storePath := "./store/"
	httpAddr := ":8080"

	db, err := database.NewJsonDataBase(storePath)
	if err != nil {
		panic(err)
	}

	taskManagerService := service.NewService(db)

	serverHandler := transport.MakeHandler(taskManagerService)

	server := &http.Server{
		Addr:    httpAddr,
		Handler: serverHandler,
	}

	server.ListenAndServe()
}
