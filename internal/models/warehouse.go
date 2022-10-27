package models

type Warehouse struct {
	// Идентификатор склада
	ID int64
	// Название склада в системе
	Name string
	// Ячейки склада
	Cells []*Cell
}

type Cell struct {
	// Идентификатор
	ID int64
	// Строковое представление адреса ячейки
	Name string
	// Номер ряда
	Row int
	// Номер секции
	Column int
	// Номер ячейки
	Num int
	// Содержимое ячейки
	Content []*Product
}
