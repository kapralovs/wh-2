package tasks

import "github.com/kapralovs/wh-2/internal/models"

type Usecase interface {
	Start()
	Complete()
}

type Repository interface {
	CreateTask() (*models.Product, error)
	DeleteTask() error
	GetTask() ([]*models.Product, error)
	GetTaskById() (*models.Product, error)
}
