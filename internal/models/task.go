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
)

// Задача
type Task struct {
	// Идентификатор задачи
	ID int64
	// Статус
	Status string
	// Тип задачи
	Event any
}

// Заказ
type Order struct {
	// Идентификатор клиента
	ClientID int64
	// Адрес доставки
	Address       string
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
}

// Движение товара
type Replacing struct {
}

// Инвентаризация
type Inventory struct {
	// Ячейка
	Cell string
	// Продукт
	Product *Product
}

// Размещение товара на складском месте
type ProductPlacement struct {
}

// Особая ситуация
type SpecialSituation struct {
	// Идентификатор особой ситуации
	ID string
}
