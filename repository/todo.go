package repository

import (
	"github.com/daniilmikhaylov2005/crudTodo/models"
	"log"
)

func GetAllTodos(userId int) []models.Todo {
	db := getConnection()
	defer db.Close()

	query := `SELECT * FROM todos WHERE userid=$1`
	rows, err := db.Query(query, userId)

	if err != nil {
		log.Fatal(err)
	}

	var todos []models.Todo

	defer rows.Close()

	for rows.Next() {
		var todo models.Todo

		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Done, &todo.UserID); err != nil {
			log.Fatal(err)
		}

		todos = append(todos, todo)
	}

	return todos

}

func CreateTodo(todo models.Todo, userId int) int {
	db := getConnection()
	defer db.Close()

	var id int

	query := `INSERT INTO todos (title, done, userid) VALUES ($1, $2, $3) RETURNING ID`
	row := db.QueryRow(query, todo.Title, todo.Done)
	row.Scan(&id)

	return id
}

func GetTodoById(id, userId int) (models.Todo, error) {
	db := getConnection()
	defer db.Close()

	var todo models.Todo

	query := `SELECT * FROM todos WHERE id=$1 AND userid=$2`
	row := db.QueryRow(query, id, userId)
	if err := row.Scan(&todo.ID, &todo.Title, &todo.Done, &todo.UserID); err != nil {
		return todo, err
	}

	return todo, nil
}

func UpdateTodo(id, userId int, todo models.Todo) error {
	db := getConnection()
	defer db.Close()

	query := `UPDATE todos SET title=$1, done=$2 WHERE id=$3 AND userid=$4`
	_, err := db.Exec(query, todo.Title, todo.Done, id, userId)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTodo(todoId, userId int) (int, error) {
	db := getConnection()
	defer db.Close()

	var id int

	query := `DELETE FROM todos WHERE id=$1 AND userid=$2 RETURNING id`
	row := db.QueryRow(query, todoId, userId)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
