package models

import "time"

// Поставка товара
type Delivery struct {
	// Идентификатор
	ID int64 `json:"id,omitempty"`
	// Идентификатор заказчика
	ClientID int64 `json:"client_id,omitempty"`
	// Ячейка размещения
	Cell string `json:"cell,omitempty"`
	// Адрес доставки
	Address string `json:"address,omitempty"`
	// Содержимое поставки
	Content []*Product `json:"content,omitempty"`
	// Дата создания
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Дата обновления
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Дата удаления
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

// Единица отгрузки
type ShippingUnit struct {
	// Идентификатор
	ID int64 `json:"id,omitempty"`
	// Идентификатор задачи
	TaskID int64 `json:"task_id,omitempty"`
	// Идентификатор исполнителя
	EmployeeID int64 `json:"employee_id,omitempty"`
	// Содержимое ЕО
	Content []*Product `json:"content,omitempty"`
	// Дата создания
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Дата обновления
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Дата удаления
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}
