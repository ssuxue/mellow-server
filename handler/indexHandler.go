package handler

import (
	"github.com/gin-gonic/gin"
	"memo/model"
	"net/http"
	"path"
	"strings"
)

func Index(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "hello gin " + strings.ToLower(context.Request.Method) + " method",
	})
}

func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		dst := path.Join("./static/images", file.Filename)
		err := c.SaveUploadedFile(file, dst)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "OK"})
		}
	}
}

func GetAll(context *gin.Context) {
	milkTea := model.MilkTea{}
	milkTeas := milkTea.FindAll()
	result := model.Result{
		Code:    http.StatusOK,
		Message: "查询成功",
		Data:    milkTeas,
	}
	context.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}


