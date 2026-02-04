package repo

import (
	"fmt"
	"golang_app_development_practice_2/internal/models"
)

func CreateTask(title string) (models.TaskResponse, error) {
	gorm := GetDB()
	newTask := models.TaskResponse{
		Title: title,
		Done:  false,
	}
	err := gorm.Create(&newTask).Error
	if err != nil {
		return models.TaskResponse{}, fmt.Errorf("create task: %s", err)
	}
	return newTask, err
}
