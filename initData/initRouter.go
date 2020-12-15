package initData

import (
	"github.com/gin-gonic/gin"
	"memo/handler"
	"net/http"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("./../templates/*")
	} else {
		router.LoadHTMLGlob("templates/*")
	}

	router.GET("/", handler.Index)
	router.GET("/getDoublePeelMilk", handler.GetAll)
	router.GET("/getMilkTeaCategory", handler.GetMilkTeaCategory)
	router.GET("/getAllMilkyTea", handler.GetMilkyTeas)
	router.GET("/getMilkyTeaByCategory/:id", handler.GetByCategory)
	router.GET("/getAllCategory", handler.GetAllCategory)
	router.GET("/getAllAttribute", handler.GetAllAttribute)
	router.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.zhihu.com")
	})
	router.POST("/upload", handler.Upload)

	userGroup := router.Group("/user")
	{
		userGroup.POST("/signup", handler.Signup)
		userGroup.POST("/login", handler.Login)
	}
	return router
}
