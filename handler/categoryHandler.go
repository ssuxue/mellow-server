package handler

import (
	"github.com/gin-gonic/gin"
	"memo/model"
	"net/http"
)

func GetMilkTeaCategory(context *gin.Context) {
	milkTeaCategory := model.MilkTeaCategory{}
	categories := milkTeaCategory.FindCategory()
	result := model.Result{
		Code:    http.StatusOK,
		Message: "查询成功",
		Data:    categories,
	}
	context.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func GetAllCategory(context *gin.Context) {
	category := model.Category{}
	categories := category.FindAll()
	result := model.Result{
		Code:    http.StatusOK,
		Message: "查询成功",
		Data:    categories,
	}
	context.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
