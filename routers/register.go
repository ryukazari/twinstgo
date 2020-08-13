package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ryukazari/twinstgo/database"
	"github.com/ryukazari/twinstgo/models"
)

// Register function to register users in the database
func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	var resp models.Response
	if err != nil {
		// http.Error(w, "Error in the received data for user: "+err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		resp.Message = "Error in the received data for user: " + err.Error()
		resp.Status = http.StatusBadRequest
		resp.Ok = false
		json.NewEncoder(w).Encode(resp)
		return
	}

	if len(t.Email) == 0 {
		// http.Error(w, "User email is required", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		resp.Message = "User email is required"
		resp.Status = http.StatusBadRequest
		resp.Ok = false
		json.NewEncoder(w).Encode(resp)
		return
	}

	if len(t.Password) < 8 {
		// http.Error(w, "Password must be at least 8 characters", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		resp.Message = "Password must be at least 8 characters"
		resp.Status = http.StatusBadRequest
		resp.Ok = false
		json.NewEncoder(w).Encode(resp)
		return
	}

	_, encontrado, _ := database.UserExists(t.Email)
	if encontrado {
		// http.Error(w, "User already exists with the email: "+t.Email, http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		resp.Message = "User already exists with the email: " + t.Email
		resp.Status = http.StatusBadRequest
		resp.Ok = false
		json.NewEncoder(w).Encode(resp)
		return
	}

	_, status, err := database.RegisterUser(t)

	if err != nil {
		// http.Error(w, "An error ocurred while registering the user"+err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		resp.Message = "An error ocurred while registering the user" + err.Error()
		resp.Status = http.StatusInternalServerError
		resp.Ok = false
		json.NewEncoder(w).Encode(resp)
		return
	}

	if !status {
		http.Error(w, "User registration failed"+err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		resp.Message = "User registration failed" + err.Error()
		resp.Status = http.StatusInternalServerError
		resp.Ok = false
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp.Message = "Succesful token"
	resp.Status = http.StatusCreated
	resp.Ok = true
	resp.Data = t
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
