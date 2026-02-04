package handlers

import (
	"encoding/json"
	"golang_app_development_practice_2/internal/repo"
	"log"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
)

func GetTask(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	idStr := strings.TrimPrefix(path, "/tasks/")
	if idStr == "" {
		WriteError(w, http.StatusBadRequest, "Missing id parameter")
		return
	}
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		WriteError(w, http.StatusBadRequest, "Invalid id format")
		return
	}
	slog.Info("GET /tasks/", id)
	task, err := repo.GetTaskById(id)
	if err != nil {
		WriteError(w, http.StatusNotFound, "Task not found")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}
