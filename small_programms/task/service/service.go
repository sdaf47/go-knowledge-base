package service

import (
	"github.com/sdaf47/go-knowledge-base/small_programms/task/db"
	"github.com/sdaf47/go-knowledge-base/small_programms/database"
)

type TaskManagerService interface {
	CreateTask(name string) (int, error)
	Tasks() ([]db.Task, error)
	Task(id int) (db.Task, error)
	DeleteTask(id int) error
	SetStatus(id int, status string) error
}

type service struct {
	taskRepo db.TaskRepository
}

func NewService(dataBase database.DataBase) TaskManagerService {
	return &service{
		taskRepo: db.NewTaskRepository(dataBase),
	}
}

func (s *service) CreateTask(name string) (int, error) {
	return s.taskRepo.Create(name)
}

func (s *service) Tasks() ([]db.Task, error) {
	return s.taskRepo.Tasks()
}

func (s *service) Task(id int) (db.Task, error) {
	return s.taskRepo.Task(id)
}

func (s *service) DeleteTask(id int) error {
	return s.taskRepo.Delete(id)
}

func (s *service) SetStatus(id int, status string) (err error) {
	task, err := s.taskRepo.Task(id)
	if err != nil {
		return
	}

	task.Status = status

	err = s.taskRepo.Update(id, task)

	return
}
