package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ryukazari/twinstgo/database"
	"github.com/ryukazari/twinstgo/models"
)

//ModifyProfile Modify user profile
func ModifyProfile(w http.ResponseWriter, r *http.Request) {
	resp := models.Response{
		Ok:     true,
		Status: http.StatusOK,
	}
	w.Header().Set("Content-Type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp.Ok = false
		resp.Status = http.StatusBadRequest
		resp.Message = "Invalid data " + err.Error()
		json.NewEncoder(w).Encode(resp)
		return
	}
	var status bool
	status, err = database.ModifyRegister(t, IDUsuario)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp.Message = "Error during modified register " + err.Error()
		resp.Status = http.StatusBadRequest
		resp.Ok = false
		json.NewEncoder(w).Encode(resp)
		return
	}
	if !status {
		w.WriteHeader(http.StatusBadRequest)
		resp.Ok = false
		resp.Status = http.StatusBadRequest
		resp.Message = "Error during modify register"
		return
	}
	w.WriteHeader(http.StatusOK)
	resp.Message = "Registry modified successfully"
	json.NewEncoder(w).Encode(resp)
	return
}
