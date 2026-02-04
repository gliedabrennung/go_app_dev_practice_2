package repo

import (
	"golang_app_development_practice_2/internal/models"
)

func GetTasks() ([]models.TaskResponse, error) {
	db := GetDB()
	var tasks []models.TaskResponse

	result := db.Table("task_responses").Select("id", "title", "done").Find(&tasks)

	if result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}
