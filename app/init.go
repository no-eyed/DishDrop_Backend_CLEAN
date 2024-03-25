package app

import (
	"database/sql"
	"meal-backend/handlers/http"
	"meal-backend/repositories"
	"meal-backend/usecases"
)

type HandlersSchema struct {
	MealHandler *http.MealHandler
	UserHandler *http.UserHandler
}

func InitialiseHandlers(db *sql.DB) *HandlersSchema {

	mealRepo := repositories.NewMealRepository(db)
	mealUsecase := usecases.NewMealUsecase(mealRepo)
	mealHandler := http.NewMealHandler(mealUsecase)

	userRepo := repositories.NewUserRepository(db)
	userUsecase := usecases.NewUserUsecase(userRepo)
	userHandler := http.NewUserHandler(userUsecase)

	handlers := &HandlersSchema{
		MealHandler: mealHandler,
		UserHandler: userHandler,
	}

	return handlers
}
