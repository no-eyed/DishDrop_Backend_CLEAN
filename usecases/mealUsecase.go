package usecases

import (
	"meal-backend/domain/entities"
)

type MealUsecase struct {
	MealRepo entities.MealRepository
}

func NewMealUsecase(repo entities.MealRepository) *MealUsecase {
	return &MealUsecase{
		MealRepo: repo,
	}
}

func (mUC *MealUsecase) GetAll() ([]*entities.ResponseMeal, error) {
	return mUC.MealRepo.GetAll()
}

func (mUC *MealUsecase) GetAllByUserId(clerkId string) ([]*entities.ResponseMeal, error) {
	userId, err := mUC.MealRepo.GetUserIdByClerkId(clerkId)
	if err != nil {
		return nil, err
	}
	return mUC.MealRepo.GetAllByUserId(userId)
}

func (mUC *MealUsecase) GetById(id int64) (*entities.Meal, error) {
	return mUC.MealRepo.GetById(id)
}

func (mUC *MealUsecase) GetResponseById(id int64) (*entities.ResponseMeal, error) {
	return mUC.MealRepo.GetResponseById(id)
}

func (mUC *MealUsecase) Save(meal *entities.Meal, clerkId string) error {
	userId, err := mUC.MealRepo.GetUserIdByClerkId(clerkId)
	if err != nil {
		return err
	}
	meal.Creator_id = userId
	return mUC.MealRepo.Save(meal)
}

func (mUC *MealUsecase) Update(meal *entities.Meal, mealId int64) error {
	meal.Id = mealId
	return mUC.MealRepo.Update(meal)
}

func (mUC *MealUsecase) Delete(mealId int64) error {
	return mUC.MealRepo.Delete(mealId)
}
