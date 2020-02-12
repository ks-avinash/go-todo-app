package Routes

import (
	"github.com/gin-gonic/gin"
	"go-todo-app/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("todo", controllers.GetTodos)
		v1.POST("todo", controllers.CreateATodo)
		v1.GET("todo/:id", controllers.GetATodo)
		v1.PUT("todo/:id", controllers.UpdateATodo)
		v1.DELETE("todo/:id", controllers.DeleteATodo)
	}

	return r
}
