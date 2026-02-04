package repo

import "golang_app_development_practice_2/internal/models"

func GetTaskById(id uint64) (*models.TaskResponse, error) {
	db := GetDB()
	var task models.TaskResponse
	if err := db.Where("id = ?", id).First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}
