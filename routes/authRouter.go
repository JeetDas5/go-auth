package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jeetdas5/golang-jwt/controllers"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("/signup", controllers.Signup())
	router.POST("/login", controllers.Login())
}
