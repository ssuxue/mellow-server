package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ssuxue/mellow-server/domain/model"
	"net/http"
)

func GetAllAttribute(context *gin.Context) {
	attribute := model.Attribute{}
	attributes := attribute.FindAll()
	result := model.Result{
		Code:    http.StatusOK,
		Message: "查询成功",
		Data:    attributes,
	}
	context.JSON(http.StatusOK, result)
}
