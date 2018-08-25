package service

import (
	"testing"
	"github.com/sdaf47/go-knowledge-base/small_programms/database"
	"os"
	"fmt"
	"time"
)

const StorePath = "./store/"

var taskManagerService TaskManagerService

func TestMain(m *testing.M) {
	db, err := database.NewJsonDataBase(StorePath)
	if err != nil {
		panic(err)
	}

	taskManagerService = NewService(db)

	os.Exit(m.Run())
}

func TestService_CreateTask(t *testing.T) {
	var err error
	var id int

	id, err = taskManagerService.CreateTask("task 1")
	if err != nil {
		t.Fatal(err)
	}

	id, err = taskManagerService.CreateTask("task 2")
	if err != nil {
		t.Fatal(err)
	}
	if id < 1 {
		t.Fatal("second id must be gritter then first")
	}
}

func TestService_Tasks(t *testing.T) {
	tasks, err := taskManagerService.Tasks()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("tasks", tasks)
}

func TestService_Task(t *testing.T) {
	task, err := taskManagerService.Task(1)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("task", task)
}

func TestService_SetStatus(t *testing.T) {
	id := 0
	status := "pending"
	err := taskManagerService.SetStatus(id, status)
	if err != nil {
		t.Fatal(err)
	}

	task, err := taskManagerService.Task(id)
	if err != nil {
		return
	}

	if task.Status != status {
		t.Fatalf("status of task not equial new status")
	}
}

func TestService_DeleteTask(t *testing.T) {
	taskName := "test_name_" + string(time.Now().Second())
	id, err := taskManagerService.CreateTask(taskName)
	if err != nil {
		t.Fatal(err)
	}

	task, err := taskManagerService.Task(id)
	if err != nil {
		t.Fatal(err)
	}

	if task.Name != taskName {
		t.Fatal("wrong task id")
	}

	err = taskManagerService.DeleteTask(id)
	if err != nil {
		t.Fatal(err)
	}

	task, err = taskManagerService.Task(id)
	if task.Name == taskName {
		t.Fatal("task was not delete")
	}
}
