package handlers

import (
	"GoATTHStart/internal/services"
	"encoding/json"
	"log/slog"
	"net/http"
)

type HealthHander struct {
	service *services.HealthService
	logger  *slog.Logger
}

func NewHealthHandler(service *services.HealthService, logger *slog.Logger) *HealthHander {
	return &HealthHander{
		service: service,
		logger:  logger,
	}
}

func (h *HealthHander) HealthHandler(w http.ResponseWriter, r *http.Request) {
	status, err := h.service.CheckHealth()
	if err != nil {
		h.logger.Error("error checking health", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	statusBytes, err := json.Marshal(status)
	if err != nil {
		h.logger.Error("error marshalling status", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(statusBytes)
}
