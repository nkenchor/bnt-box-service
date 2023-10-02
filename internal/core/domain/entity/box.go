package entity

import (
	"bnt/bnt-box-service/internal/core/domain/shared"
	"time"
)

type Box struct {
	Reference    string        `json:"reference" bson:"reference"`
	BatchNo      string        `json:"reference_no" bson:"batch_no"`
	BoxNo        string         `json:"box_no" bson:"box_no"`
	Prefix       string        `json:"prefix" bson:"prefix"`
	Indent       Indent        `json:"indent" bson:"indent"`
	Status       shared.Status `json:"status" bson:"status"`
	Transactions []Transaction `json:"transactions" bson:"transactions"`
	Movement     Movement      `json:"movement" bson:"movement"`
	CreatedAt    time.Time     `json:"created_at" bson:"created_at"`
	CreatedBy    User          `json:"created_by" bson:"created_by"`
	UpdatedAt    time.Time     `json:"updated_at" bson:"updated_at"`
	UpdatedBy    User          `json:"updated_by" bson:"updated_by"`
}
