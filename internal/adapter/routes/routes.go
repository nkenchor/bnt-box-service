package routes

import (
	"bnt/bnt-box-service/internal/adapter/api"
	errorhelper "bnt/bnt-box-service/internal/core/helper/error-helper"
	logger "bnt/bnt-box-service/internal/core/helper/log-helper"
	message "bnt/bnt-box-service/internal/core/helper/message-helper"
	"bnt/bnt-box-service/internal/core/middleware"
	"bnt/bnt-box-service/internal/core/services"
	ports "bnt/bnt-box-service/internal/port"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	boxRepository ports.BoxRepository) *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	boxService := services.NewBox(boxRepository)

	handler := api.NewHTTPHandler(
		boxService)

	logger.LogEvent("INFO", "Configuring Routes!")
	router.Use(middleware.LogRequest)

	//router.Use(middleware.SetHeaders)

	router.Group("/api/box")
	{
		router.POST("/api/box/generate", handler.Generate)
		router.GET("/api/box/indent/:reference", handler.GetIndentByRef)

	}


	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404,
			errorhelper.ErrorMessage(errorhelper.NoResourceError, message.NoResourceFound))
	})

	return router
}
