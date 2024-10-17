package controller

import (
	"log/slog"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte(`{"message": "API cloud run success"}`))
	if err != nil {
		slog.Error("failed to write response", "Error", err.Error())
	}
}
