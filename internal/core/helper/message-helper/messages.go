package helper

import (
	configuration "bnt/bnt-box-service/internal/core/helper/configuration-helper"
	"fmt"
	"time"
)

var (
	NoResourceFound = "this resource does not exist"
	NoRecordFound   = "sorry. no record found"
	ServerStarted   = fmt.Sprintf("Started " + configuration.GlobalConfig.ServiceName + " on " + configuration.GlobalConfig.ServiceAddress + ":" +
		configuration.GlobalConfig.ServicePort + " in " + time.Since(time.Now()).String())
)

var (
	ValidationError = "VALIDATION_ERROR"
	RedisSetupError = "REDIS_SETUP_ERROR"
	NoRecordError   = "NO_RECORD_FOUND_ERROR"
	NoResourceError = "INVALID_RESOURCE_ERROR"
	CreateError     = "CREATE_ERROR"
	UpdateError     = "UPDATE_ERROR"
	LogError        = "LOG_ERROR"
	MongoDBError    = "MONGO_DB_ERROR"
)
