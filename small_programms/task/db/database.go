package db

import (
	"time"
	"github.com/sdaf47/go-knowledge-base/small_programms/database"
)

type Task struct {
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
	Status  string    `json:"status"`
}

type TaskRepository interface {
	Task(id int) (Task, error)
	Tasks() ([]Task, error)
	Create(name string) (id int, err error)
	Delete(id int) error
	Update(id int, task Task) error
}

type taskRepository struct {
	db database.DataBase
}

func NewTaskRepository(db database.DataBase) TaskRepository {
	return &taskRepository{
		db: db,
	}
}

func (tr *taskRepository) Task(id int) (t Task, err error) {
	res, err := tr.db.GetOne(id, Task{})
	if err != nil {
		return
	}

	t = res.(Task)

	return
}

func (tr *taskRepository) Tasks() (tasks []Task, err error) {
	res, err := tr.db.Get(Task{})
	if err != nil {
		return
	}

	tasks = res.([]Task)

	return
}

func (tr *taskRepository) Create(name string) (id int, err error) {
	id, err = tr.db.Create(Task{
		Name:    name,
		Created: time.Now(),
	})

	return
}

func (tr *taskRepository) Delete(id int) (err error) {
	err = tr.db.Delete(id, Task{})

	return
}

func (tr *taskRepository) Update(id int, task Task) (err error) {
	err = tr.db.Update(id, task)

	return
}
