package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ryukazari/twinstgo/database"
	"github.com/ryukazari/twinstgo/models"
)

//ExistRelation  endpoint to verify if a relation exists
func ExistRelation(w http.ResponseWriter, r *http.Request) {
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
	t.UserFollowedID = ID
	t.UserID = IDUsuario

	var respuesta models.RelationResponse

	status, err := database.ExistRelation(t)

	if err != nil || !status {
		w.WriteHeader(http.StatusOK)
		respuesta.Status = false
		resp.Ok = false
		resp.Message = "Relation doesn't exists"
		resp.Status = http.StatusBadRequest
		json.NewEncoder(w).Encode(respuesta)
		return
	}
	respuesta.Status = true

	w.WriteHeader(http.StatusOK)
	resp.Message = "Relation exists!"
	resp.Data = respuesta
	json.NewEncoder(w).Encode(resp)
	return
}
