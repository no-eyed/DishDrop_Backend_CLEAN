package routes

import (
	"meal-backend/app"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, handlers *app.HandlersSchema) {
	SetMealRoutes(router, handlers)
	SetUserRoutes(router, handlers)
}
