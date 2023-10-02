package entity

type Denomination struct {
	Denomination string `json:"denomination"`
	Value        int    `json:"value"`
	Code         string `json:"code"`
}
