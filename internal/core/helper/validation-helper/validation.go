package helper

import (
	errorhelper "bnt/bnt-box-service/internal/core/helper/error-helper"
	logger "bnt/bnt-box-service/internal/core/helper/log-helper"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func Validate(data interface{}) error {
	logger.LogEvent("INFO", "Validating "+reflect.TypeOf(data).String()+" Data...")
	err := validator.New().Struct(data)
	if err != nil {
		var fieldErrors []validator.FieldError
		logger.LogEvent("ERROR", "Error validating struct: "+err.Error())
		for _, errs := range err.(validator.ValidationErrors) {
			fieldErrors = append(fieldErrors, errs)
		}
		return errorhelper.ErrorArrayToError(fieldErrors)
	}
	logger.LogEvent("INFO", reflect.TypeOf(data).String()+" Data Validated Successfully...")
	return nil
}
