package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if payload == nil {
		return
	}

	_ = json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{
		"error": message,
	})
}

func parseIDParam(r *http.Request, param string) (int, error) {
	value := chi.URLParam(r, param)
	if value == "" {
		return 0, fmt.Errorf("missing path param %s", param)
	}

	id, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("invalid %s", param)
	}

	return id, nil
}
