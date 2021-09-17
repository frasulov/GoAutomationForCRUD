package models

type Model struct {
	Name string
	Fields []FieldsArray
}

type FieldsArray struct {
	Column Column
}

type Column struct {
	Name string
	Type string
}
