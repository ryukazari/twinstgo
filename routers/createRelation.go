package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ryukazari/twinstgo/database"

	"github.com/ryukazari/twinstgo/models"
)

//CreateRelation endpoint to create a relation between two users
func CreateRelation(w http.ResponseWriter, r *http.Request) {
	resp := models.Response{
		Ok:     true,
		Status: http.StatusOK,
	}
	w.Header().Set("Content-Type", "application/json")

	ID := r.URL.Query().Get("id") // ID User followed
	if len(ID) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		resp.Ok = false
		resp.Message = "Must to send ID parameter"
		resp.Status = http.StatusBadRequest
		json.NewEncoder(w).Encode(resp)
		return
	}

	var t models.Relation

	t.UserID = IDUsuario
	t.UserFollowedID = ID

	status, err := database.CreateRelation(t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp.Ok = false
		resp.Message = "An error ocurred while create relation " + err.Error()
		resp.Status = http.StatusInternalServerError
		json.NewEncoder(w).Encode(resp)
		return
	}
	if !status {
		w.WriteHeader(http.StatusInternalServerError)
		resp.Ok = false
		resp.Message = "An error ocurred while insert relation in BD"
		resp.Status = http.StatusInternalServerError
		json.NewEncoder(w).Encode(resp)
		return
	}
	w.WriteHeader(http.StatusOK)
	resp.Message = "Relation created succesfully!"
	json.NewEncoder(w).Encode(resp)
	return
}
