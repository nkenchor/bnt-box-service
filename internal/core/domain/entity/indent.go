package entity


type Indent struct {
	Reference    string       `json:"reference" bson:"reference"`
	Location     Location     `json:"location" bson:"location" validate:"required"`
	Year         Year         `json:"year" bson:"year" validate:"required"`
	Denomination Denomination `json:"denomination" bson:"denomination" validate:"required,min=1"`
}
