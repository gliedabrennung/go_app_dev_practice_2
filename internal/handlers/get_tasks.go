package handlers

import (
	"encoding/json"
	"golang_app_development_practice_2/internal/models"
	"golang_app_development_practice_2/internal/repo"
	"log/slog"
	"net/http"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	slog.Info("GET /tasks ", r)
	tasks, err := repo.GetTasks()
	if err != nil {
		slog.Error("GetTasks error:", err)
		WriteError(w, http.StatusInternalServerError, err.Error())
	}
	if tasks == nil {
		tasks = []models.TaskResponse{}
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		return
	}
}
