package repository

import (
	"github.com/daniilmikhaylov2005/crudTodo/models"
	"log"
)

func GetAllTodos() []models.Todo {
	db := getConnection()
	defer db.Close()

	query := `SELECT * FROM todos`
	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	var todos []models.Todo

	defer rows.Close()

	for rows.Next() {
		var todo models.Todo

		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Done); err != nil {
			log.Fatal(err)
		}

		todos = append(todos, todo)
	}

	return todos

}

func CreateTodo(todo models.Todo) int {
	db := getConnection()
	defer db.Close()

	var id int

	query := `INSERT INTO todos (title, done) VALUES ($1, $2) RETURNING ID`
	row := db.QueryRow(query, todo.Title, todo.Done)
	row.Scan(&id)

	return id
}

func GetTodoById(id int) (models.Todo, error) {
	db := getConnection()
	defer db.Close()

	var todo models.Todo

	query := `SELECT * FROM todos WHERE id=$1`
	row := db.QueryRow(query, id)
	if err := row.Scan(&todo.ID, &todo.Title, &todo.Done); err != nil {
		return todo, err
	}

	return todo, nil
}

func UpdateTodo(id int, todo models.Todo) error {
	db := getConnection()
	defer db.Close()

	query := `UPDATE todos SET title=$1, done=$2 WHERE id=$3`
	_, err := db.Exec(query, todo.Title, todo.Done, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTodo(id int) (int, error) {
	db := getConnection()
	defer db.Close()

	query := `DELETE FROM todos WHERE id=$1`
	_, err := db.Exec(query, id)
	if err != nil {
		return 0, nil
	}
	return id, nil
}
