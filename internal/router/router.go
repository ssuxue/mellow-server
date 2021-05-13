package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ssuxue/mellow-server/handler"
	"net/http"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/getAllMilkyTea", handler.GetMilkyTeas)
	router.GET("/getBubbleTeaByCid/:id", handler.GetByCategory)
	router.GET("/getAllCategory", handler.GetAllCategory)
	router.GET("/getAllAttribute", handler.GetAllAttribute)
	router.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.zhihu.com")
	})

	userGroup := router.Group("/user")
	{
		userGroup.POST("/signup", handler.Signup)
		userGroup.POST("/login", handler.Login)
	}
	return router
}
