package routers

import (
	"golang_app_development_practice_2/internal/handlers"
	"golang_app_development_practice_2/internal/middleware"
	"net/http"
)

func InitRouter() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("GET /tasks", handlers.GetTasks)
	router.HandleFunc("POST /tasks", handlers.PostTasks)
	router.HandleFunc("PATCH /tasks", handlers.UpdateTasks)
	router.HandleFunc("GET /tasks/", handlers.GetTask)
	return middleware.Auth(router)
}
