package http

import (
	"meal-backend/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MealHandler struct {
	MealUsecase entities.MealUsecase
}

func NewMealHandler(uc entities.MealUsecase) *MealHandler {
	return &MealHandler{
		MealUsecase: uc,
	}
}

func (mealHandler *MealHandler) GetAllMeals(context *gin.Context) {
	meals, err := mealHandler.MealUsecase.GetAll()
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, meals)
}

func (mealHandler *MealHandler) GetMealById(context *gin.Context) {
	mealId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse meal id"})
		return
	}

	meal, err := mealHandler.MealUsecase.GetResponseById(mealId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch meal", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, meal)
}

func (mealHandler *MealHandler) GetAllMealsByUserId(context *gin.Context) {
	clerkId := context.Param("clerkid")
	meals, err := mealHandler.MealUsecase.GetAllByUserId(clerkId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch meals", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, meals)
}

func (mealHandler *MealHandler) AddMeal(context *gin.Context) {
	meal := &entities.Meal{}
	err := context.BindJSON(&meal)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request body"})
	}

	clerkId := context.Param("clerkid")

	err = mealHandler.MealUsecase.Save(meal, clerkId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save meal", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, meal)
}

func (mealHandler *MealHandler) UpdateMeal(context *gin.Context) {
	meal := &entities.Meal{}
	err := context.BindJSON(&meal)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request body"})
		return
	}

	mealId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse meal id"})
		return
	}

	err = mealHandler.MealUsecase.Update(meal, mealId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update meal", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, meal)
}

func (mealHandler *MealHandler) DeleteMeal(context *gin.Context) {
	mealId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse meal id"})
	}

	err = mealHandler.MealUsecase.Delete(mealId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete meal", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, mealId)
}
