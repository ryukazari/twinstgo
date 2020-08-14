package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ryukazari/twinstgo/database"

	"github.com/ryukazari/twinstgo/models"
)

//DeleteTwit execute deleteTwit that delete a twit from database
func DeleteTwit(w http.ResponseWriter, r *http.Request) {
	resp := models.Response{
		Ok:     true,
		Status: http.StatusOK,
	}
	w.Header().Set("Content-Type", "application/json")
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		resp.Ok = false
		resp.Message = "You must pass parameter ID"
		resp.Status = http.StatusBadRequest
		json.NewEncoder(w).Encode(resp)
		return
	}

	err := database.DeleteTwit(ID, IDUsuario)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp.Ok = false
		resp.Message = "An error ocurred while delete a twit"
		resp.Status = http.StatusInternalServerError
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(http.StatusOK)
	resp.Message = "Twit deleted successfully."
	json.NewEncoder(w).Encode(resp)
	return
}
