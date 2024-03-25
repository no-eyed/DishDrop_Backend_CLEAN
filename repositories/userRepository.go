package repositories

import (
	"database/sql"
	"meal-backend/domain/entities"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (u *UserRepository) GetUserById(id int64) (*entities.User, error) {
	query := `SELECT * FROM users WHERE id = $1`

	row := u.DB.QueryRow(query, id)

	var user *entities.User

	err := row.Scan(&user.Id, &user.ClerkId, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) GetUserIdByClerkId(clerkId string) (int64, error) {
	query := `SELECT id FROM users WHERE clerk_id = $1`

	row := u.DB.QueryRow(query, clerkId)

	var id int64

	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *UserRepository) Save(user *entities.User) error {
	query := `INSERT INTO users (clerk_id, username, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	stmt, err := u.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.ClerkId, user.Username, user.Email, user.CreatedAt, user.UpdatedAt)
	err = result.Scan(&user.Id)

	return err
}

func (u *UserRepository) Update(user *entities.User) error {
	query := `UPDATE users SET username = $1, email = $2, updated_at = $3 WHERE id = $4`

	stmt, err := u.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Email, user.UpdatedAt, user.Id)

	return err
}

func (u *UserRepository) Delete(id int64) error {
	query := `DELETE FROM users WHERE id = $1`

	stmt, err := u.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	return err
}
