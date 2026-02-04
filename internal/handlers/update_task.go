package handlers

import (
	"encoding/json"
	"errors"
	"golang_app_development_practice_2/internal/repo"
	"log"
	"log/slog"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

func UpdateTasks(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	idStr := query.Get("id")
	if idStr == "" {
		WriteError(w, http.StatusBadRequest, "Missing id parameter")
		return
	}
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		WriteError(w, http.StatusBadRequest, "Invalid id format")
		return
	}
	slog.Info("PATCH /tasks/", id)
	task, err := repo.UpdateTask(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		WriteError(w, http.StatusNotFound, "Task not found")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}
