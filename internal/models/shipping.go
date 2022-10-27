package models

// Поставка товара
type Delivery struct {
	// Идентификатор
	ID int64
	// Идентификатор заказчика
	UserID int64
	// Ячейка размещения
	Cell string
	// Содержимое поставки
	Content []*Product
}

// Единица отгрузки
type ShippingUnit struct {
	// Идентификатор
	ID int64
	// Идентификатор задачи
	TaskID int64
	// Идентификатор исполнителя
	UserID int64
	// Содержимое ЕО
	Content []*Product
}
