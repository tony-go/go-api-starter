package routes

import "github.com/gin-gonic/gin"

func ApplyRoutes(r *gin.Engine) {
	root := r.Group("/")
	{
		ApplyUserRoutes(root)
	}
}
