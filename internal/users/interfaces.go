package users

import "github.com/kapralovs/wh-2/internal/models"

type UseCase interface {
	CreateUser() (*models.User, error)
	EditID(id int64) error
	GetUsers() ([]*models.User, error)
	GetUserByID(id int64) (*models.User, error)
	DeleteUser(id int64) error
}

type Repository interface {
	CreateUser() (*models.User, error)
	EditID(id int64) error
	GetUsers() ([]*models.User, error)
	GetUserByID(id int64) (*models.User, error)
	DeleteUser(id int64) error
}
