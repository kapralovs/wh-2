package models

// Поставка товара
type Delivery struct {
	// Идентификатор
	ID int64
	// Содержимое поставки
	Content []*Product
}

// Единица отгрузки
type ShippingUnit struct {
	// Идентификатор
	ID int64
	// Идентификатор исполнителя
	UserID int64
	// Содержимое ЕО
	Content []*Product
}
