package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ryukazari/twinstgo/models"
)

// Register function to register users in the database
func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error in the received data for user: "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "User email is required", http.StatusBadRequest)
		return
	}

	if len(t.Password) < 8 {
		http.Error(w, "Password must be at least 8 characters", http.StatusBadRequest)
		return
	}

	_, encontrado, _ := database.UserExists(t.Email)
	if encontrado {
		http.Error(w, "User already exists with the email"+t.Email, http.StatusBadRequest)
		return
	}

	_, status, err := database.RegisterUser(t)

	if err != nil {
		http.Error(w, "An error ocurred while registering the user"+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(w, "User registration failed"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
