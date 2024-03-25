package routes

import (
	"meal-backend/app"
	"meal-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func SetMealRoutes(router *gin.Engine, handlers *app.HandlersSchema) {
	router.GET("/meals", handlers.MealHandler.GetAllMeals)
	router.GET("/meals/:id", handlers.MealHandler.GetMealById)
	router.GET("/meals/my-meals/:clerkid", middlewares.VerifyUser(), handlers.MealHandler.GetAllMealsByUserId)
	router.POST("/meals/:clerkid", middlewares.AuthMiddleware(), handlers.MealHandler.AddMeal)
	router.PUT("/meals/:clerkid/:id", middlewares.VerifyUser(), handlers.MealHandler.UpdateMeal)
	router.DELETE("/meals/:clerkid/:id", middlewares.VerifyUser(), handlers.MealHandler.DeleteMeal)
}
