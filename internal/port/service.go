package ports

import (
	"bnt/bnt-box-service/internal/core/domain/dto"
)

type BoxService interface {
	Generate(generationConfig dto.BoxConfig) (interface{}, error)
	GetIndentByRef(CountryCode string) (interface{}, error)
}
