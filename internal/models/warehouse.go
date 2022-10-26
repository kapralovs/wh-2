package models

type Cell struct {
	ID      string
	Content []*Product
}

type Row struct {
	ID      string
	Columns []*Column
}

type Column struct {
	ID    string
	Cells []*Cell
}
