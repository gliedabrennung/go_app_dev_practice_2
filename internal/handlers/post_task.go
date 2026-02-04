package handlers

import (
	"encoding/json"
	"golang_app_development_practice_2/internal/models"
	"golang_app_development_practice_2/internal/repo"
	"log"
	"log/slog"
	"net/http"
)

func PostTasks(w http.ResponseWriter, r *http.Request) {
	req, err := decodeJSON[models.TaskResponse](r)
	if err != nil {
		WriteError(w, http.StatusBadRequest, "Invalid JSON or unknown field payload received")
		return
	}

	slog.Info("POST /tasks ", req)
	task, err := repo.CreateTask(req.Title)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}
