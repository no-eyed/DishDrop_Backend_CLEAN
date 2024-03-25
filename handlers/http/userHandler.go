package http

import (
	"meal-backend/domain/entities"
	"meal-backend/domain/requests"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUsecase entities.UserUsecase
}

func NewUserHandler(usecase entities.UserUsecase) *UserHandler {
	return &UserHandler{
		UserUsecase: usecase,
	}
}

func (uH *UserHandler) GetUser(context *gin.Context) {
	clerkId := context.Param("clerkId")
	userId, err := uH.UserUsecase.GetUserIdByClerkId(clerkId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uH.UserUsecase.GetUserById(userId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, user)
}

func (uH *UserHandler) UserManager(context *gin.Context) {
	var user entities.User

	var req requests.UserRequest

	err := context.BindJSON(&req)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ClerkId = req.Data.UUID

	if req.Type == "user.deleted" {

		err := uH.UserUsecase.Delete(&user)

		if err != nil {

			context.JSON(http.StatusInternalServerError, gin.H{
				"message": "could not delete user",
			})

			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": "user deleted successfully",
		})

		return
	}

	user.Email = req.Data.Email[0].EmailAddress
	user.Username = req.Data.UserName
	user.CreatedAt = time.Unix(req.Data.CreatedAt/1000, 0)
	user.UpdatedAt = time.Unix(req.Data.UpdatedAt/1000, 0)

	if req.Type == "user.updated" {

		err = uH.UserUsecase.Update(&user)

		if err != nil {

			context.JSON(http.StatusInternalServerError, gin.H{
				"message": "could not update user",
			})

			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": "user updated successfully",
		})

		return

	} else {
		err = uH.UserUsecase.Save(&user)

		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"message": "could not create user",
			})

			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": "user created successfully",
		})

	}
}

func (uH *UserHandler) CreateUser(context *gin.Context) {
	var user entities.User
	err := context.BindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = uH.UserUsecase.Save(&user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, user)
}

func (uH *UserHandler) UpdateUser(context *gin.Context) {
	var user entities.User
	err := context.BindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = uH.UserUsecase.Update(&user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, user)
}

func (uH *UserHandler) DeleteUser(context *gin.Context) {
	var user entities.User
	err := context.BindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = uH.UserUsecase.Delete(&user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, user)
}
