package main

import (
	"leaflet-back/internal/config"
	"leaflet-back/internal/handlers"
	"leaflet-back/internal/services"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	log.Info("Logger initialized")

	cfg := config.MustLoad()
	log.Info("Config loaded", slog.String("env", cfg.Env))

	mbtilesService := services.NewMbtilesService(cfg.StoragePath)
	log.Info("Services initialized")

	handler := handlers.NewHandler(mbtilesService)
	log.Info("Handlers initialized")

	log.Info("Starting server", slog.String("address", cfg.Address))

	server := &http.Server{
		Addr:         cfg.Address,
		Handler:      handler.InitRoutes(),
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Error("Failed to start server", "error", err)
	}
}
