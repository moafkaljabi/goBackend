package server

import (
	"encoding/json"
	"goBackend/internal/models"
	"net/http"
	"strconv"
)

func (s *APIServer) handleCreateDevice(w http.ResponseWriter, r *http.Request) error {
	var device models.Device
	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return err
	}

	err := s.store.CreateDevice(&device)
	if err != nil {
		http.Error(w, "Failed to create device", http.StatusInternalServerError)
		return err
	}

	json.NewEncoder(w).Encode(device)
	return WriteJSON(w, http.StatusCreated, device)
}

func (s *APIServer) handleGetDevicesByUserID(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	devices, err := s.store.GetDevicesByUserID(id)
	if err != nil {
		http.Error(w, "Failed to fetch devices", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(devices)
}

func (s *APIServer) handleGetDeviceByID(w http.ResponseWriter, r *http.Request) {

}
