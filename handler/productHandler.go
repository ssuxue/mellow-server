package handler

import (
	"github.com/gin-gonic/gin"
	"memo/model"
	"net/http"
	"strconv"
)

func GetMilkyTeas(context *gin.Context) {
	milkyTea := model.Product{}
	categories := milkyTea.FindAll()

	result := model.Result{
		Code:    http.StatusOK,
		Message: "查询成功",
		Data:    categories,
	}

	context.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func GetByCategory(context *gin.Context) {
	milkyTea := model.Product{}
	id, _ := strconv.Atoi(context.Param("id"))
	categories := milkyTea.FindByCategoryId(id)

	result := model.Result{
		Code:    http.StatusOK,
		Message: "查询成功",
		Data:    categories,
	}

	context.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
