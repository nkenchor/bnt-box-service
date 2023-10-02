package api

import (
	"bnt/bnt-box-service/internal/core/domain/dto"
	errorhelper "bnt/bnt-box-service/internal/core/helper/error-helper"
	validation "bnt/bnt-box-service/internal/core/helper/validation-helper"

	"github.com/gin-gonic/gin"
)

func (hdl *HTTPHandler) Generate(c *gin.Context) {
	body := dto.BoxConfig{}
	_ = c.BindJSON(&body)

	if err := validation.Validate(&body); err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}

	box, err := hdl.boxService.Generate(body)
	if err != nil {
		c.AbortWithStatusJSON(500, err)
		return
	}
	c.JSON(201, gin.H{"reference": box})
}

func (hdl *HTTPHandler) GetIndentByRef(c *gin.Context) {
	indent, err := hdl.boxService.GetIndentByRef(c.Param("reference"))

	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}

	c.JSON(200, indent)
}
