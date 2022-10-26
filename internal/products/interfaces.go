package products

import "github.com/kapralovs/wh-2/internal/models"

type Usecase interface {
}

type Repository interface {
	CreateProduct() (*models.Product, error)
	DeleteProduct() error
	GetProducts() ([]*models.Product, error)
	GetProductById() (*models.Product, error)
}
