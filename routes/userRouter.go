package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jeetdas5/golang-jwt/controllers"
	"github.com/jeetdas5/golang-jwt/middleware"
)

func UserRoutes(router *gin.Engine) {
	userRoutes := router.Group("/")
	userRoutes.Use(middleware.AuthMiddleware())
	{
		userRoutes.GET("/users", controllers.GetUsers())
		userRoutes.GET("/users/:user_id", controllers.GetUser())
	}
}
