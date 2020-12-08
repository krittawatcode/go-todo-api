package repository

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // use to connect db
	"github.com/krittawatcode/go-todo-api/config"
	"github.com/krittawatcode/go-todo-api/models"
)

// GetAllTodos use for GET ALL todo list
func GetAllTodos(todo *[]models.Todo) (err error) {
	if err = config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

// CreateATodo use for CREATE todo
func CreateATodo(todo *models.Todo) (err error) {
	if err = config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

// GetATodo use for GET specific Todo by ID
func GetATodo(todo *models.Todo, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

// UpdateATodo use for UPDATE specific Todo by ID
func UpdateATodo(todo *models.Todo, id string) (err error) {
	fmt.Println(todo)
	config.DB.Save(todo) // save all field
	return nil
}

// DeleteATodo use for DELETE specific Todo by ID
func DeleteATodo(todo *models.Todo, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(todo)
	return nil
}
