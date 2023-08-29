package handlers

import (
	"encoding/json"
	"forum/internal/models"
	"forum/internal/services"
	"net/http"
	"strconv"
	"time"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var credentials models.Credentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	// Authenticate the user
	email, password := credentials.Email, credentials.Password
	user, err := h.authService.Login(email, password)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

<<<<<<< HEAD
	currentSession, _ := session.Storee.Get(r, user.Username)

	currentSession.Values["username"] = user.Username
	currentSession.Values["userEmail"] = user.Email
	currentSession.Values["userID"] = user.Id

=======
>>>>>>> master
	cookie := http.Cookie{
		Name:     "sessionID",   // Replace "myCookie" with the desired name for your cookie
		Value:    user.Username, // Replace "cookieValue123" with the desired value for your cookie
		HttpOnly: true,
		Path:     "/", // Optional: set the cookie's path
		// You can add more optional settings for the cookie, like Domain, Expires, MaxAge, HttpOnly, Secure, etc.
	}

	cookie2 := http.Cookie{
		Name:     "userID",              // Replace "myCookie" with the desired name for your cookie
		Value:    strconv.Itoa(user.Id), // Replace "cookieValue123" with the desired value for your cookie
		HttpOnly: true,
		Path:     "/", // Optional: set the cookie's path
		// You can add more optional settings for the cookie, like Domain, Expires, MaxAge, HttpOnly, Secure, etc.
	}

	http.SetCookie(w, &cookie)

	http.SetCookie(w, &cookie2)

	// Return the authenticated user in the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode user data", http.StatusInternalServerError)
		return
	}
}


func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// Clear the cookie by setting an expired cookie

	http.SetCookie(w, &http.Cookie{
		Name:     "sessionID",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour), // Set the expiration time to a past value
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1, // Make sure to set the same path used when setting the cookie
	})

<<<<<<< HEAD
	sessionKey := sessionID.Value

	sessionInfos, _ := session.Storee.Get(r, sessionKey)

	for key := range sessionInfos.Values {
		delete(sessionInfos.Values, key)
	}
=======
	http.SetCookie(w, &http.Cookie{
		Name:     "userID",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour), // Set the expiration time to a past value
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1, // Make sure to set the same path used when setting the cookie
	})
>>>>>>> master

	// Return success message in the response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User logged out successfully"))
}


func (h *AuthHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var resetRequest models.ResetRequest
	if err := json.NewDecoder(r.Body).Decode(&resetRequest); err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	// Reset the password
	err := h.authService.ResetPassword(resetRequest.Email, resetRequest.NewPassword)
	if err != nil {
		http.Error(w, "Failed to reset password", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Password reset successfully"))
}