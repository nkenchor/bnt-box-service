package ports

import "bnt/bnt-box-service/internal/core/domain/entity"

type BoxRepository interface {
	Generate(boxes []entity.Box) (interface{}, error)
	GetIndentByRef(CountryCode string) (interface{}, error)
}
