package repository

import (
	"github.com/daniilmikhaylov2005/crudTodo/models"
)

func InsertUser(user models.User) (string, error) {
	db := getConnection()
	defer db.Close()

	var id string

	query := `INSERT INTO users (username, name, role, password) VALUES ($1, $2, $3, $4) RETURNING id`

	row := db.QueryRow(query, user.Username, user.Name, user.Role, user.Password)
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func FindUserByUsername(username string) (models.User, error) {
	return models.User{}, nil
}
