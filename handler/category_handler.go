package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ssuxue/mellow-server/domain/model"
	"net/http"
)

func GetAllCategory(context *gin.Context) {
	category := model.Category{}
	categories := category.FindAll()

	result := model.Result{
		Code:    http.StatusOK,
		Message: "查询成功",
		Data:    categories,
	}

	//context.JSON(http.StatusOK, gin.H{
	//	"result": result,
	//})
	context.JSON(http.StatusOK, result)
}
