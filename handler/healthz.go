package handler

import (
	"net/http"

	"encoding/json"

	"github.com/TechBowl-japan/go-stations/model"

	"log"
)

// A HealthzHandler implements health check endpoint.
type HealthzHandler struct{}

// NewHealthzHandler returns HealthzHandler based http.Handler.
func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{}
}

// ServeHTTP implements http.Handler interface.
func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var healthRes = model.HealthzResponse{Message: "OK"}
	err := json.NewEncoder(w).Encode(healthRes)
	if err != nil {
		log.Println(err)
	}
}
