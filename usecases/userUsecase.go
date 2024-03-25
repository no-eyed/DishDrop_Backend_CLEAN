package usecases

import "meal-backend/domain/entities"

type UserUsecase struct {
	UserRepo entities.UserRepository
}

func NewUserUsecase(repo entities.UserRepository) *UserUsecase {
	return &UserUsecase{
		UserRepo: repo,
	}
}

func (uUC *UserUsecase) GetUserById(id int64) (*entities.User, error) {
	return uUC.UserRepo.GetUserById(id)
}

func (uUC *UserUsecase) GetUserIdByClerkId(clerkId string) (int64, error) {
	return uUC.UserRepo.GetUserIdByClerkId(clerkId)
}

func (uUC *UserUsecase) Save(user *entities.User) error {
	return uUC.UserRepo.Save(user)
}

func (uUC *UserUsecase) Update(user *entities.User) error {
	userId, err := uUC.UserRepo.GetUserIdByClerkId(user.ClerkId)
	if err != nil {
		return err
	}
	user.Id = userId

	return uUC.UserRepo.Update(user)
}

func (uUC *UserUsecase) Delete(user *entities.User) error {
	userId, err := uUC.UserRepo.GetUserIdByClerkId(user.ClerkId)
	if err != nil {
		return err
	}
	return uUC.UserRepo.Delete(userId)
}
