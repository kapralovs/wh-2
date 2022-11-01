package tasks

import "github.com/kapralovs/wh-2/internal/models"

type UseCase interface {
	CreateTask() (*models.Product, error)
	DeleteTask(id int64) error
	GetTask() ([]*models.Product, error)
	GetTaskById(id int64) (*models.Product, error)
	EditTask(id int64) error
	AssignTo(userID int64) error
	Start()
	Complete()
}

type Repository interface {
	CreateTask() (*models.Product, error)
	DeleteTask(id int64) error
	GetTask() ([]*models.Product, error)
	GetTaskById(id int64) (*models.Product, error)
	EditTask(id int64) error
}
