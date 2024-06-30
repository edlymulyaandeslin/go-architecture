package repository

import (
	"clean-code-app-laundry/model"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

// buat interface
type UserRepository interface {
	GetAll(page int, size int) ([]model.User, error)
	GetById(id string) (model.User, error)
	CreateUser(payload model.User) (model.User, error)
	FindByUsername(username string) (model.User, error)
	isUsernameExists(username string) bool
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

func (p *userRepository) CreateUser(payload model.User) (model.User, error) {
	u := model.User{}
	err := p.db.QueryRow(`INSERT INTO users (name, email, username, password, role, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, name, email, username, role, created_at`, payload.Name, payload.Email, payload.Username, payload.Password, payload.Role, time.Now()).Scan(
		&u.Id,
		&u.Name,
		&u.Email,
		&u.Username,
		&u.Role,
		&u.CreatedAt,
	)

	if err != nil {
		return model.User{}, err
	}

	return u, nil
}

func (p *userRepository) FindByUsername(username string) (model.User, error) {

	if !p.isUsernameExists(username) {
		err := fmt.Sprintf("user dengan username %v tidak ditemukan!", username)
		return model.User{}, errors.New(err)
	}

	user := model.User{}
	err := p.db.QueryRow(`SELECT id, name, email, username, password, role, created_at, updated_at FROM users WHERE username=$1`, username).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (p *userRepository) isUsernameExists(username string) bool {
	usernameList := []string{}
	rows, _ := p.db.Query(`SELECT username FROM users`)
	for rows.Next() {
		var username string
		rows.Scan(&username)
		usernameList = append(usernameList, username)
	}

	var isExists bool
	for _, v := range usernameList {
		if username == v {
			isExists = true
			break
		} else {
			isExists = false
		}
	}
	return isExists
}

// constructor
func NewUserRepository(database *sql.DB) UserRepository {
	return &userRepository{db: database}
}
