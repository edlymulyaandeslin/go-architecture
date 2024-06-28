package repository

import (
	"clean-code-app-laundry/model"
	"database/sql"
)

// buat interface
type UserRepository interface {
	GetAll(page int, size int) ([]model.User, error)
	GetById(id string) (model.User, error)
}

// struct
type userRepository struct {
	db *sql.DB
}

func (p *userRepository) GetAll(page int, size int) ([]model.User, error) {
	panic("unimplemented")
}

func (p *userRepository) GetById(id string) (model.User, error) {
	var user model.User

	err := p.db.QueryRow("SELECT id, name, email, username, password, role, created_at, updated_at FROM users WHERE id =$1", id).Scan(&user.Id, &user.Name, &user.Email, &user.Username, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

// constructor
func NewUserRepository(database *sql.DB) UserRepository {
	return &userRepository{db: database}
}
