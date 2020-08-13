package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ryukazari/twinstgo/database"
	"github.com/ryukazari/twinstgo/models"
)

// ViewProfile extract values from profile
func ViewProfile(w http.ResponseWriter, r *http.Request) {
	resp := models.Response{
		Ok:     true,
		Status: http.StatusOK,
	}

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		// http.Error(w, "Must to send ID parameter", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		resp.Ok = false
		resp.Message = "Must to send ID parameter"
		resp.Status = http.StatusBadRequest
		json.NewEncoder(w).Encode(resp)
		return
	}

	perfil, err := database.SearchPerfil(ID)
	if err != nil {
		// http.Error(w, "An error ocurred during search profile: "+err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		resp.Ok = false
		resp.Message = "An error ocurred during search profile: " + err.Error()
		resp.Status = http.StatusBadRequest
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp.Data = perfil
	json.NewEncoder(w).Encode(resp)
}
