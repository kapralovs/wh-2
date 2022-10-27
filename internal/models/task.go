package models

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
	ID int64
	// Тип задачи
	TypeOf string
	// На кого присвоена задача
	AssignedTo int64
	// Статус
	Status string
	// Содержимое задачи
	Event any
}

// Заказ
type Order struct {
	// Идентификатор
	ID int64
	// Идентификатор поставки
	DeliveryID int64
	// Идентификатор клиента
	ClientID int64
	// Адрес доставки
	Address string
	// Единицы отгрузки
	ShippingUnits []*ShippingUnit
}

// Пополнение пик-зоны
type Refill struct {
	// Отпускающее складское место
	ReleasingCell string
	// Товар
	ProductID int64
	// Кол-во единиц товара
	Count int64
	// Принимающее складское место
	HostCell string
}

// Отгрузка
type Transport struct {
	// Идентификатор
	ID int64
	// Отпускающий склад
	ReleasingWarehouse string
	// Идентификатор поставки
	DeliveryID int64
	// Идентификатор единицы отгрузки
	ShippingUnitID int64
}

// Движение товара
type Replacing struct {
	// Идентификатор
	ID int64
	// Отпускающее складское место
	ReleasingCell string
	// Товар
	ProductID int64
	// Кол-во единиц товара
	ProductCount int
	// Принимающее складское место
	HostCell string
}

// Инвентаризация
type Inventory struct {
	// Идентификатор
	ID int64
	// Ячейка
	Cell string
	// Продукт
	ProductID int64
	// Заявленное кол-во
	DeclaredQuantity int
	// Фактическое кол-во
	ActualQuantity int
}

// Размещение товара на складском месте
type ProductPlacement struct {
	// Идентификатор
	ID int64
	// Отпускающее складское место
	ReleasingCell string
	// Товар
	ProductID int64
	// Кол-во единиц товара
	ProductCount int
	// Принимающее складское место
	HostCell string
}

// Особая ситуация
type SpecialSituation struct {
	// Идентификатор
	ID int64
	// Название
	Name string
	// Описание
	Description string
}
