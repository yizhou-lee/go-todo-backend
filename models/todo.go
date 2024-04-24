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

func GetTodos() []Todo {
	var todos []Todo
	sql.GetDB().Find(&todos)
	return todos
}

func CreateTodo(t *Todo) *Todo {
	sql.GetDB().Create(t)
	return t
}

func GetTodoByID(id int) *Todo {
	var t Todo
	sql.GetDB().First(&t, id)
	return &t
}

func UpdateTodoByID(id int, t *Todo) *Todo {
	var todo Todo
	sql.GetDB().First(&todo, id)
	if todo.ID != 0 {
		sql.GetDB().Model(&todo).Updates(t)
	}
	return &todo
}

func DeleteTodoByID(id int) *Todo {
	var t Todo
	sql.GetDB().First(&t, id)
	if t.ID != 0 {
		sql.GetDB().Delete(&t)
	}
	return &t
}
