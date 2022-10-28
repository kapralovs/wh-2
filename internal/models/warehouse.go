package models

import "time"

type Warehouse struct {
	// Идентификатор склада
	ID int64 `json:"id,omitempty"`
	// Название склада в системе
	Name string `json:"name,omitempty"`
	// Ячейки склада
	Cells []*Cell `json:"cells,omitempty"`
	// Дата создания
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Дата обновления
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Дата удаления
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

type Cell struct {
	// Идентификатор
	ID int64 `json:"id,omitempty"`
	// Строковое представление адреса ячейки
	Name string `json:"name,omitempty"`
	// Номер ряда
	Row int `json:"row,omitempty"`
	// Номер секции
	Column int `json:"column,omitempty"`
	// Номер ячейки
	Num int `json:"num,omitempty"`
	// Содержимое ячейки
	Content []*Product `json:"content,omitempty"`
	// Дата создания
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Дата обновления
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Дата удаления
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}
