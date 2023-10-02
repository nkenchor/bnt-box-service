package helper

import (
	helper "bnt/bnt-box-service/internal/core/helper/configuration-helper"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var (
	ValidationError = "VALIDATION_ERROR"
	RedisSetupError = "REDIS_SETUP_ERROR"
	NoRecordError   = "NO_RECORD_FOUND_ERROR"
	NoResourceError = "INVALID_RESOURCE_ERROR"
	CreateError     = "CREATE_ERROR"
	UpdateError     = "UPDATE_ERROR"
	LogError     = "LOG_ERROR"
	MongoDBError = "MONGO_DB_ERROR"
)




func (err ErrorResponse) Error() string {
	var errorBody ErrorBody
	return fmt.Sprintf("%v", errorBody)

}
func ErrorArrayToError(errorBody []validator.FieldError) error {
	var errorResponse ErrorResponse
	errorResponse.TimeStamp = time.Now().Format(time.RFC3339)
	errorResponse.ErrorReference = uuid.New()
    errorResponse.ErrorType = ValidationError
	for _, value := range errorBody {
		body := ErrorBody{Message: value.Error()}
		errorResponse.Errors = append(errorResponse.Errors, body)
	}
	return errorResponse
}
func ErrorMessage(errorType string, message string) error {
	var errorResponse ErrorResponse
	errorResponse.TimeStamp = time.Now().Format(time.RFC3339)
	errorResponse.ErrorReference = uuid.New()
	errorResponse.ErrorType = errorType
	errorResponse.ErrorSource = helper.GlobalConfig.ServiceName
	errorResponse.Errors = append(errorResponse.Errors, ErrorBody{Message: message})
	return errorResponse
}

type ErrorBody struct {
	//Code    string      `json:"code"`
	Message interface{} `json:"message"`
	//Source  string      `json:"source"`
}
type ErrorResponse struct {
	ErrorReference uuid.UUID   `json:"error_reference"`
	ErrorType      string 		`json:"error_type"`
	TimeStamp      string      `json:"timestamp"`
	ErrorSource	   string 		`json:"error_source"`
	Errors         []ErrorBody `json:"errors"`
}