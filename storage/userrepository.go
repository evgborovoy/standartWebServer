package storage

import (
	"fmt"
	"log"

	"github.com/evgborovoy/StandardWebServer/internal/app/models"
)

type UserRepository struct {
	storage *Storage
}

var (
	tableUser string = "users"
)

// Create user
func (ur *UserRepository) Create(user *models.User) (*models.User, error) {
	query := fmt.Sprintf("INSERT INTO %s (login, password) VALUES($1, $2) RETURNING id", tableUser)
	if err := ur.storage.db.QueryRow(query, user.Login, user.Password).Scan(&user.ID); err != nil {
		return nil, err
	}
	return user, nil
}

// Find user
func (ur *UserRepository) FindByLogin(login string) (*models.User, bool, error) {
	founded := false
	users, err := ur.SelectAll()
	if err != nil {
		return nil, founded, err
	}
	var userFinded *models.User
	for _, u := range users {
		if u.Login == login {
			userFinded = u
			founded = true
			break
		}
	}
	return userFinded, founded, nil
}

// Select all users
func (ur *UserRepository) SelectAll() ([]*models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableUser)
	rows, err := ur.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := make([]*models.User, 0)
	for rows.Next() {
		u := models.User{}
		err := rows.Scan(&u.ID, &u.Login, &u.Password)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, &u)
	}
	return users, nil
}
