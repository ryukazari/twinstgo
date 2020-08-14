package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ryukazari/twinstgo/database"
	"github.com/ryukazari/twinstgo/models"
)

//CreateTwit create a twit
func CreateTwit(w http.ResponseWriter, r *http.Request) {
	resp := models.Response{
		Ok:     true,
		Status: http.StatusOK,
	}
	w.Header().Set("Content-Type", "application/json")
	var message models.Twit
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp.Ok = false
		resp.Message = "Invalid data " + err.Error()
		resp.Status = http.StatusBadRequest
		json.NewEncoder(w).Encode(resp)
		return
	}
	registro := models.CreateTwit{
		UserID:  IDUsuario,
		Message: message.Message,
		Date:    time.Now(),
	}
	var status bool
	_, status, err = database.CreateTwit(registro)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp.Ok = false
		resp.Message = "An error occurred while trying to register the twit" + err.Error()
		resp.Status = http.StatusBadRequest
		json.NewEncoder(w).Encode(resp)
		return
	}

	if !status {
		w.WriteHeader(http.StatusBadRequest)
		resp.Ok = false
		resp.Message = "The tweet could not be registered"
		resp.Status = http.StatusBadRequest
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp.Message = "Twit registered successful"
	json.NewEncoder(w).Encode(resp)
	return

}
