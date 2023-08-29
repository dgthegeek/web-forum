package handlers

import (
	"encoding/json"
	"forum/internal/models"
	"forum/internal/services"
	"log"
	"net/http"
)

type RegistrationHandler struct {
	registrationService services.RegistrationService
}

func NewRegistrationHandler(registrationService services.RegistrationService) *RegistrationHandler {
	return &RegistrationHandler{
		registrationService: registrationService,
	}
}

func (h *RegistrationHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Parse the request body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Failed to parse form data: %v", err)
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	// Register the user using the registration service
	err = h.registrationService.RegisterUser(&user)
	if err != nil {
		log.Printf("Failed to register user: %v", err)
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Printf("Failed to encode user data: %v", err)
		http.Error(w, "Failed to encode user data", http.StatusInternalServerError)
		return
	}
}
