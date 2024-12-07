package service

import (
	"go-todo/config"
	"go-todo/models"
)

// Get all todos
func GetAllTodos() ([]models.Todo, error) {
	var todos []models.Todo
	err := config.DB.Find(&todos).Error
	if err != nil {
		return nil, err
	}

	return todos, nil
}

// Get a todo by id
func GetTodoById(id int) (models.Todo, error) {
	var todo models.Todo
	err := config.DB.First(&todo, id).Error
	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

// Create a todo
func CreateTodo(input models.Todo) (models.Todo, error) {
	err := config.DB.Create(&input).Error
	if err != nil {
		return models.Todo{}, err
	}

	return input, nil
}

// Update a todo by id
func UpdateTodoById(id int, input models.Todo) (models.Todo, error) {
	todo, err := GetTodoById(id)
	if err != nil {
		return models.Todo{}, err
	}

	err = config.DB.Model(&todo).Updates(input).Error
	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

// Delete a todo by id
func DeleteTodoById(id int) (models.Todo, error) {
	todo, err := GetTodoById(id)
	if err != nil {
		return models.Todo{}, err
	}

	err = config.DB.Delete(&todo).Error
	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}