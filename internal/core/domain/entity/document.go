package entity

type Document struct {
	Reference       string     	`json:"reference" bson:"reference"`
	Filename        string      `json:"filename" bson:"filename"`
}