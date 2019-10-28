package routes

import (
	"api/handlers"
	"github.com/gin-gonic/gin"
)

func ApplyUserRoutes (r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.GET("/", handlers.GetUsers)
		user.GET("/:id", handlers.GetUser)
		user.POST("/", handlers.CreateUser)
		user.DELETE("/:id", handlers.DeleteUser)
		user.PATCH("/:id", handlers.PatchUser)
	}
}
