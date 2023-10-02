package services

import (
	"bnt/bnt-box-service/internal/core/domain/dto"
	"bnt/bnt-box-service/internal/core/domain/entity"

	//"bnt/bnt-box-service/internal/core/domain/mapper"
	helper "bnt/bnt-box-service/internal/core/helper/application-helper"
	logger "bnt/bnt-box-service/internal/core/helper/log-helper"
	ports "bnt/bnt-box-service/internal/port"
)

type boxService struct {
	boxRepository ports.BoxRepository
}

func NewBox(boxRepository ports.BoxRepository) *boxService {
	return &boxService{
		boxRepository: boxRepository,
	}
}

func (service *boxService) Generate(generationConfig dto.BoxConfig) (interface{}, error) {

	indent, _ := service.GetIndentByRef(generationConfig.IndentRef)
	boxlist, err := helper.Generate(generationConfig.PrefixStart,generationConfig.PrefixEnd,generationConfig.PrefixType, indent.(entity.Indent))
	if err != nil {
		return nil, err
	}
	boxes, err := service.boxRepository.Generate(boxlist.([]entity.Box))
	if err != nil {
		return nil, err
	}
	return boxes, nil
}

func (service *boxService) GetIndentByRef(reference string) (interface{}, error) {
	logger.LogEvent("INFO", "Getting indent with reference: "+reference)
	indent, err := service.boxRepository.GetIndentByRef(reference)

	if err != nil {
		return nil, err
	}
	return indent, nil
}
