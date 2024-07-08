package models

import (
	"todo-backend/pkg/mysql"

	"gorm.io/gorm"
)

// Todo represents a todo item
//	@Description	A single todo item
//	@Name			Todo
type Todo struct {
	gorm.Model
	Title       string `json:"title" example:"Buy milk"`
	Completed   bool   `json:"completed" example:"false"`
	Description string `json:"description" example:"Remember to buy milk from the grocery store"`
}

func GetTodos() []Todo {
	var todos []Todo
	mysql.GetDB().Find(&todos)
	return todos
}

func CreateTodo(t *Todo) *Todo {
	mysql.GetDB().Create(t)
	return t
}

func GetTodoByID(id int) *Todo {
	var t Todo
	mysql.GetDB().First(&t, id)
	return &t
}

func UpdateTodoByID(id int, t *Todo) *Todo {
	var todo Todo
	mysql.GetDB().First(&todo, id)
	if todo.ID != 0 {
		mysql.GetDB().Model(&todo).Updates(t)
	}
	return &todo
}

func DeleteTodoByID(id int) *Todo {
	var t Todo
	mysql.GetDB().First(&t, id)
	if t.ID != 0 {
		mysql.GetDB().Delete(&t)
	}
	return &t
}
