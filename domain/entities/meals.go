package entities

type Meal struct {
	Id           int64  `json:"id"`
	Title        string `json:"title"`
	Image        string `json:"image"`
	Summary      string `json:"summary"`
	Instructions string `json:"instructions"`
	Creator_id   int64  `json:"creator_id"`
}

type ResponseMeal struct {
	Id            int64  `json:"id"`
	Title         string `json:"title"`
	Image         string `json:"image"`
	Summary       string `json:"summary"`
	Instructions  string `json:"instructions"`
	Creator_name  string `json:"creator"`
	Creator_email string `json:"creator_email"`
}

type MealRepository interface {
	GetAll() ([]*ResponseMeal, error)
	GetUserIdByClerkId(clerkId string) (int64, error)
	GetAllByUserId(id int64) ([]*ResponseMeal, error)
	GetById(id int64) (*Meal, error)
	GetResponseById(id int64) (*ResponseMeal, error)
	Save(meal *Meal) error
	Update(meal *Meal) error
	Delete(mealId int64) error
}

type MealUsecase interface {
	GetAll() ([]*ResponseMeal, error)
	GetAllByUserId(clerkId string) ([]*ResponseMeal, error)
	GetById(id int64) (*Meal, error)
	GetResponseById(id int64) (*ResponseMeal, error)
	Save(meal *Meal, clerkId string) error
	Update(meal *Meal, mealId int64) error
	Delete(mealId int64) error
}
