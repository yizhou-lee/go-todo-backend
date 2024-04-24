package models

import (
	"todo-backend/pkg/sql"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string `json:"title"`
	Completed   bool   `json:"completed"`
	Description string `json:"description"`
}

var db = sql.GetDB()

func GetTodos() []Todo {
	var todos []Todo
	db.Find(&todos)
	return todos
}

func CreateTodo(t *Todo) *Todo {
	db.Create(t)
	return t
}

func GetTodoByID(id int) *Todo {
	var t Todo
	db.First(&t, id)
	return &t
}

func UpdateTodoByID(id int, t *Todo) *Todo {
	var todo Todo
	db.First(&todo, id)
	if todo.ID != 0 {
		db.Model(&todo).Updates(t)
	}
	return &todo
}

func DeleteTodoByID(id int) *Todo {
	var t Todo
	db.First(&t, id)
	if t.ID != 0 {
		db.Delete(&t)
	}
	return &t
}
