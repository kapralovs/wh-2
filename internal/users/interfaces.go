package users

import "github.com/kapralovs/wh-2/internal/models"

type Usecase interface {
	CreateUser() (*models.User, error)
	DeleteUser() error
}

type Repository interface {
	CreateUser() (*models.User, error)
	DeleteUser() error
	GetUsers() ([]*models.User, error)
	GetUserById() (*models.User, error)
}
