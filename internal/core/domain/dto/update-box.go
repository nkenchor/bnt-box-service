package dto

import (
	entity "bnt/bnt-box-service/internal/core/domain/entity"
)

type UpdateBox struct {
	Transactions []entity.Transaction `json:"transactions" bson:"transactions"`
}
