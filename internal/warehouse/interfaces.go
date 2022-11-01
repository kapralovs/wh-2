package warehouse

import "github.com/kapralovs/wh-2/internal/models"

type UseCase interface {
	CreateWarehouse() (*models.Warehouse, error)
	EditWarehouse(id int64) error
	GetWarehouses() ([]*models.Warehouse, error)
	GetWarehouseByID(id int64) (*models.Warehouse, error)
	DeleteWarehouse(id int64) error
}

type Repository interface {
	CreateWarehouse() (*models.Warehouse, error)
	EditWarehouse(id int64) error
	GetWarehouses() ([]*models.Warehouse, error)
	GetWarehouseByID(id int64) (*models.Warehouse, error)
	DeleteWarehouse(id int64) error
}
