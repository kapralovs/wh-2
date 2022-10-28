package models

import "time"

const (
	BIDU = "BIDU"
	BIDF = "BIDF"

	StatusUndone     = "Undone"
	StatusInProgress = "In progress"
	StatusDone       = "Done"

	InboundQueue  = "Inbound"
	OutboundQueue = "Outbound"
	InternQueue   = "Intern"
	TransQueue    = "Trans"

	OrderTask            = "order"
	RefillTask           = "refill"
	TransportTask        = "transport"
	Replacingtask        = "replacing"
	InventoryTask        = "inventory"
	ProductPlacementTask = "product_placement"
)

// Задача
type Task struct {
	// Идентификатор задачи
	ID int64 `json:"id,omitempty"`
	// Тип задачи
	TypeOf string `json:"type_of,omitempty"`
	// На кого присвоена задача
	AssignedTo int64 `json:"assigned_to,omitempty"`
	// Статус
	Status string `json:"status,omitempty"`
	// Содержимое задачи
	Event any `json:"event,omitempty"`
	// Дата создания
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Дата обновления
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Дата удаления
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

// Заказ
type Order struct {
	// Идентификатор
	ID int64 `json:"id,omitempty"`
	// Идентификатор поставки
	DeliveryID int64 `json:"delivery_id,omitempty"`
	// Идентификатор клиента
	ClientID int64 `json:"client_id,omitempty"`
	// Адрес доставки
	Address string `json:"address,omitempty"`
	// Единицы отгрузки
	ShippingUnits []*ShippingUnit `json:"shipping_units,omitempty"`
}

// Пополнение пик-зоны
type Refill struct {
	// Отпускающее складское место
	ReleasingCell string `json:"releasing_cell,omitempty"`
	// Товар
	ProductID int64 `json:"product_id,omitempty"`
	// Кол-во единиц товара
	Count int64 `json:"count,omitempty"`
	// Принимающее складское место
	HostCell string `json:"host_cell,omitempty"`
}

// Отгрузка
type Transport struct {
	// Идентификатор
	ID int64 `json:"id,omitempty"`
	// Отпускающий склад
	ReleasingWarehouse string `json:"releasing_warehouse,omitempty"`
	// Идентификатор поставки
	DeliveryID int64 `json:"delivery_id,omitempty"`
	// Идентификатор единицы отгрузки
	ShippingUnitID int64 `json:"shipping_unit_id,omitempty"`
}

// Движение товара
type Replacing struct {
	// Идентификатор
	ID int64 `json:"id,omitempty"`
	// Отпускающее складское место
	ReleasingCell string `json:"releasing_cell,omitempty"`
	// Товар
	ProductID int64 `json:"product_id,omitempty"`
	// Кол-во единиц товара
	ProductCount int `json:"product_count,omitempty"`
	// Принимающее складское место
	HostCell string `json:"host_cell,omitempty"`
}

// Инвентаризация
type Inventory struct {
	// Идентификатор
	ID int64 `json:"id,omitempty"`
	// Ячейка
	Cell string `json:"cell,omitempty"`
	// Продукт
	ProductID int64 `json:"product_id,omitempty"`
	// Заявленное кол-во
	DeclaredQuantity int `json:"declared_quantity,omitempty"`
	// Фактическое кол-во
	ActualQuantity int `json:"actual_quantity,omitempty"`
}

// Размещение товара на складском месте
type ProductPlacement struct {
	// Идентификатор
	ID int64 `json:"id,omitempty"`
	// Отпускающее складское место
	ReleasingCell string `json:"releasing_cell,omitempty"`
	// Товар
	ProductID int64 `json:"product_id,omitempty"`
	// Кол-во единиц товара
	ProductCount int `json:"product_count,omitempty"`
	// Принимающее складское место
	HostCell string `json:"host_cell,omitempty"`
}

// Особая ситуация
type SpecialSituation struct {
	// Идентификатор
	ID int64 `json:"id,omitempty"`
	// Название
	Name string `json:"name,omitempty"`
	// Описание
	Description string `json:"description,omitempty"`
}
