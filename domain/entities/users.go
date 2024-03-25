package entities

import "time"

type User struct {
	Id        int64     `json:"id"`
	ClerkId   string    `json:"clerk_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository interface {
	GetUserById(id int64) (*User, error)
	GetUserIdByClerkId(clerkId string) (int64, error)
	Save(user *User) error
	Update(user *User) error
	Delete(id int64) error
}

type UserUsecase interface {
	GetUserById(id int64) (*User, error)
	GetUserIdByClerkId(clerkId string) (int64, error)
	Save(user *User) error
	Update(user *User) error
	Delete(user *User) error
}
