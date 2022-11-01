package products

import (
	"github.com/kapralovs/wh-2/internal/models"
)

type UseCase interface {
	CreateProduct() (*models.Product, error)
	EditProduct(id int64) error
	GetProducts() ([]*models.Product, error)
	GetProductById(id int64) (*models.Product, error)
	DeleteProduct(id int64) error
}

type Repository interface {
	CreateProduct() (*models.Product, error)
	EditProduct(id int64) error
	GetProducts() ([]*models.Product, error)
	GetProductById(id int64) (*models.Product, error)
	DeleteProduct(id int64) error
}
