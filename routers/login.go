package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ryukazari/twinstgo/database"
	"github.com/ryukazari/twinstgo/jwt"

	"github.com/ryukazari/twinstgo/models"
)

// Login login the user
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.User
	resp := models.LoginResponse{
		Ok:     true,
		Status: http.StatusOK,
	}

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		// http.Error(w, "Incorrect username and/or password: "+err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		resp.Ok = false
		resp.Message = "Incorrect username and/or password: " + err.Error()
		resp.Status = http.StatusBadRequest
		json.NewEncoder(w).Encode(resp)
		return
	}
	if len(t.Email) == 0 {
		// http.Error(w, "User email is required", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		resp.Ok = false
		resp.Message = "User email is required"
		resp.Status = http.StatusBadRequest
		json.NewEncoder(w).Encode(resp)
		return
	}
	document, exists := database.TryLogin(t.Email, t.Password)
	if !exists {
		// http.Error(w, "Incorrect username and/or password", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		resp.Ok = false
		resp.Message = "Incorrect username and/or password"
		resp.Status = http.StatusBadRequest
		json.NewEncoder(w).Encode(resp)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		// http.Error(w, "An error ocurred while trying to generate the token. Error: "+err.Error(), http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resp.Ok = false
		resp.Message = "An error ocurred while trying to generate the token. Error: " + err.Error()
		resp.Status = http.StatusInternalServerError
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp.Message = "Succesful token"
	resp.Token = jwtKey

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

	/*Set cookie*/
	/*
		expirationTime := time.Now().Add(20 * time.Hour)
		http.SetCookie(w, &http.Cookie{
			Name: "token_twinstgo",
			Value: jwtKey,
			Expires: expirationTime,
		})
	*/

}
