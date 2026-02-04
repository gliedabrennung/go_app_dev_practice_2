package repo

import (
	"golang_app_development_practice_2/internal/models"
)

func UpdateTask(id uint64) (models.TaskResponse, error) {
	db := GetDB()
	var task models.TaskResponse

	if err := db.First(&task, id).Error; err != nil {
		return models.TaskResponse{}, err
	}

	if err := db.Model(&task).Update("done", true).Error; err != nil {
		return models.TaskResponse{}, err
	}

	return task, nil
}
