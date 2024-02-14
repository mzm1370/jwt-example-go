package handlers

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", LoginHandler)
	}

	user := router.Group("/user")
	{
		user.POST("/register", RegisterHandler)
	}
	admin := router.Group("/admin")
	{
		admin.GET("/welcome", WelcomeHandler())
	}

}
