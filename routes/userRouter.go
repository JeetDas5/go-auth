package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jeetdas5/golang-jwt/controllers"
	"github.com/jeetdas5/golang-jwt/middleware"
)

func UserRoutes(router *gin.Engine) {
	router.Use(middleware.AuthMiddleware())
	router.GET("/users", controllers.GetUsers())
	router.GET("/users/:user_id", controllers.GetUser())
}
