package main

import (
	"golang_app_development_practice_2/internal/repo"
	"golang_app_development_practice_2/internal/routers"
	"golang_app_development_practice_2/internal/shutdown"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	err := repo.InitDB()
	if err != nil {
		slog.Error("Could not open db", err)
	}
	r := routers.InitRouter()
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	slog.Info("Starting server")
	go func() {
		err = server.ListenAndServe()
		if err != nil {
			slog.Error("Server could not start:", err)
		}
	}()
	defer func() {
		shutdown.Shutdown(server)
		err := repo.CloseDB()
		if err != nil {
			log.Println(err)
		}
	}()
}
