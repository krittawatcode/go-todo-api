package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/krittawatcode/go-todo-api/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("todo", func(c *gin.Context) {
			controllers.TDC = controllers.TODOHandler()
			controllers.TDC.GetToDos(c)
		})
		v1.POST("todo", func(c *gin.Context) {
			controllers.TDC = controllers.TODOHandler()
			controllers.TDC.CreateATodo(c)
		})
		v1.GET("todo/:id", func(c *gin.Context) {
			controllers.TDC = controllers.TODOHandler()
			controllers.TDC.GetATodo(c)
		})
		v1.PUT("todo/:id", func(c *gin.Context) {
			controllers.TDC = controllers.TODOHandler()
			controllers.TDC.UpdateATodo(c)
		})
		v1.DELETE("todo/:id", func(c *gin.Context) {
			controllers.TDC = controllers.TODOHandler()
			controllers.TDC.DeleteATodo(c)
		})
	}
	return r
}
