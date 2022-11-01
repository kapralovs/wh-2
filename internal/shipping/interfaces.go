package shipping

import "github.com/kapralovs/wh-2/internal/models"

type UseCase interface {
	CreateDelivery() (*models.Delivery, error)
	EditDelivery(id int64) error
	GetDeliveries() ([]*models.Delivery, error)
	GetDeliveryByID(id int64) (*models.Delivery, error)
	DeleteDelivery(id int64) error
}

type Repository interface {
	CreateDelivery() (*models.Delivery, error)
	EditDelivery(id int64) error
	GetDeliveries() ([]*models.Delivery, error)
	GetDeliveryByID(id int64) (*models.Delivery, error)
	DeleteDelivery(id int64) error
}
