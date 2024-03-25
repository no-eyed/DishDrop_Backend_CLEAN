package routes

import (
	"meal-backend/app"

	"github.com/gin-gonic/gin"
)

func SetUserRoutes(router *gin.Engine, handlers *app.HandlersSchema) {
	router.GET("/user/:clerkId", handlers.UserHandler.GetUser)
	router.POST("/user", handlers.UserHandler.UserManager)
	router.PUT("/user/:clerkId", handlers.UserHandler.UpdateUser)
	router.DELETE("/user/:clerkId", handlers.UserHandler.DeleteUser)
}
