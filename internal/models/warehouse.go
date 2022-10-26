package models

type Warehouse struct {
	ID    int64
	Cells []*Cell
}

type Cell struct {
	ID      string
	Content []*Product
}
