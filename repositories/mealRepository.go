package repositories

import (
	"database/sql"
	"meal-backend/domain/entities"
)

type MealRepository struct {
	DB *sql.DB
}

func NewMealRepository(db *sql.DB) *MealRepository {
	return &MealRepository{
		DB: db,
	}
}

func (m *MealRepository) GetAll() ([]*entities.ResponseMeal, error) {
	rows, err := m.DB.Query(`SELECT meals.id, title, image, summary, instructions, username, email  FROM meals INNER JOIN users ON meals.creator_id = users.id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var meals []*entities.ResponseMeal

	for rows.Next() {
		var meal entities.ResponseMeal
		err := rows.Scan(&meal.Id, &meal.Title, &meal.Image, &meal.Summary, &meal.Instructions, &meal.Creator_name, &meal.Creator_email)
		if err != nil {
			return nil, err
		}
		meals = append(meals, &meal)
	}

	return meals, nil
}

func (m *MealRepository) GetUserIdByClerkId(clerkId string) (int64, error) {
	query := `SELECT id FROM users WHERE clerk_id = $1`

	row := m.DB.QueryRow(query, clerkId)

	var id int64

	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *MealRepository) GetAllByUserId(id int64) ([]*entities.ResponseMeal, error) {
	query := `SELECT meals.id, title,image, summary, instructions, username, email  FROM meals INNER JOIN users ON meals.creator_id = users.id WHERE meals.creator_id = $1`

	rows, err := m.DB.Query(query, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var meals []*entities.ResponseMeal

	for rows.Next() {
		var meal entities.ResponseMeal
		err := rows.Scan(&meal.Id, &meal.Title, &meal.Image, &meal.Summary, &meal.Instructions, &meal.Creator_name, &meal.Creator_email)
		if err != nil {
			return nil, err
		}
		meals = append(meals, &meal)
	}

	return meals, nil
}

func (m *MealRepository) GetById(id int64) (*entities.Meal, error) {
	query := `SELECT * FROM meals WHERE id = $1`
	row := m.DB.QueryRow(query, id)

	var meal entities.Meal

	if err := row.Scan(&meal.Id, &meal.Title, &meal.Image, &meal.Summary, &meal.Instructions, &meal.Creator_id); err != nil {
		return nil, err
	}

	return &meal, nil
}

func (m *MealRepository) GetResponseById(id int64) (*entities.ResponseMeal, error) {
	query := `SELECT meals.id, title, image, summary, instructions, username, email  FROM meals INNER JOIN users ON meals.creator_id = users.id WHERE meals.id = $1`

	stmt, err := m.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var meal entities.ResponseMeal

	result := stmt.QueryRow(id)
	err = result.Scan(&meal.Id, &meal.Title, &meal.Image, &meal.Summary, &meal.Instructions, &meal.Creator_name, &meal.Creator_email)

	if err != nil {
		return nil, err
	}

	return &meal, nil
}

func (m *MealRepository) Save(meal *entities.Meal) error {
	query := `INSERT INTO meals (title, image, summary, instructions, creator_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	stmt, err := m.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result := stmt.QueryRow(meal.Title, meal.Image, meal.Summary, meal.Instructions, meal.Creator_id)
	err = result.Scan(&meal.Id)

	return err
}

func (m *MealRepository) Update(meal *entities.Meal) error {
	query := `UPDATE meals SET title = $1, image = $2, summary = $3, instructions = $4 WHERE id = $5`

	stmt, err := m.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(meal.Title, meal.Image, meal.Summary, meal.Instructions, meal.Id)

	if err != nil {
		return err
	}

	return err
}

func (m *MealRepository) Delete(id int64) error {
	query := `DELETE FROM meals WHERE id = $1`

	stmt, err := m.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	return err
}
