package mapper

import (
	"bnt/bnt-box-service/internal/core/domain/entity"
	"bnt/bnt-box-service/internal/core/domain/shared"
	
	"time"

	"github.com/google/uuid"
)

func MapBoxConfigDto(prefix string, indent entity.Indent) entity.Box {

	return entity.Box{
		BatchNo: uuid.New().String(),
		Prefix: indent.Denomination.Code + prefix,
		Indent: indent,
		Status: shared.Normal,
		Transactions: []entity.Transaction{},
		Movement: entity.Movement{},
		CreatedAt: time.Now().UTC(),
		CreatedBy: entity.User{},
		UpdatedAt: time.Now().UTC(),
		UpdatedBy: entity.User{},
	  }

}
