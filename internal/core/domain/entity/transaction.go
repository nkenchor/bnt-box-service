package entity

import (
	"bnt/bnt-box-service/internal/core/domain/shared"
	"time"
)

type Transaction struct {
	TransactionBatchNo string        `json:"transaction_batch_no" bson:"transaction_batch_no"`
	Status             shared.Status `json:"status" bson:"status"`
	TransactionBy      User          `json:"transaction_by" bson:"transaction_by"`
	TransactionTime    time.Time     `json:"transaction_time" bson:"transaction_time"`
}
