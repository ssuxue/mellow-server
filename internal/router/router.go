package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ssuxue/mellow-server/handler"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/getAllMilkyTea", handler.GetMilkyTeas)
	router.GET("/getBubbleTeaByCid/:id", handler.GetByCategory)
	router.GET("/getAllCategory", handler.GetAllCategory)
	router.GET("/getAllAttribute", handler.GetAllAttribute)
	userGroup := router.Group("/user")
	{
		userGroup.POST("/signup", handler.Signup)
		userGroup.POST("/login", handler.Login)
	}
	return router
}
