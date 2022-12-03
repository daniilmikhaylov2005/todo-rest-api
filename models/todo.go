package models

type Todo struct {
	ID    int    `json:"id"`
	UserID int `json:"userId"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}
