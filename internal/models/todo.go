package models

import (
	"gorm.io/gorm/clause"
	"todo-backend/pkg/mysql"
)

// Todo represents a todo item
//
//	@Description	A single todo item
//	@Name			Todo
type Todo struct {
	ID          int    `json:"id" gorm:"primaryKey" example:"1"`
	Title       string `json:"title" example:"Buy milk"`
	Completed   bool   `json:"completed" example:"false"`
	Description string `json:"description" example:"Remember to buy milk from the grocery store"`
}

func CreateTodo(t *Todo) error {
	if err := mysql.GetDB().Create(t).Error; err != nil {
		return err
	}
	return nil
}

func GetTodos() ([]Todo, error) {
	var todos []Todo
	if err := mysql.GetDB().Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func GetTodoByID(id int) (*Todo, error) {
	var t Todo
	if err := mysql.GetDB().First(&t, id).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func UpdateTodoByID(id int, t *Todo) error {
	var todo Todo
	tx := mysql.GetDB().Begin()

	if err := tx.First(&todo, id).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&todo).Updates(t).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func DeleteTodoByID(id int) error {
	var t Todo
	tx := mysql.GetDB().Begin()

	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&t, id).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Delete(&t).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}
