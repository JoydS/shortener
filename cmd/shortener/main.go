package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"shortener/database/model"
	"shortener/database/postgres"
	"shortener/router"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func loadEnvFiles() {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Error("Failed to load env files", "Error", err.Error())
	}
}

func main() {
	loadEnvFiles()

	db := postgres.GetGormDB()
	err := db.AutoMigrate(&model.Shortener{})
	if err != nil {
		slog.Error("Failed to migrate database", "Error", err.Error())

		return
	}

	server := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: router.Shortener(),
	}

	logger := slog.With("Port", os.Getenv("PORT"))

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		logger.Info("Shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			logger.Error("Error shutting down server", "Error", err.Error())
		}

		if err := postgres.CloseGormDB(); err != nil {
			logger.Error("Error closing database", "Error", err.Error())
		}
	}()

	logger.Info("Server is ready to handle requests")
	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		logger.Error("Error listening", "Error", err.Error())
	}

	slog.Info("Server stopped")
}
