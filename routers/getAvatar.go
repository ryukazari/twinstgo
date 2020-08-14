package routers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/ryukazari/twinstgo/database"

	"github.com/ryukazari/twinstgo/models"
)

//GetAvatar getting avatar from directory upload
func GetAvatar(w http.ResponseWriter, r *http.Request) {
	resp := models.Response{
		Ok:     true,
		Status: http.StatusOK,
	}
	w.Header().Set("Content-Type", "application/json")

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		resp.Ok = false
		resp.Message = "Must to send ID parameter"
		resp.Status = http.StatusBadRequest
		json.NewEncoder(w).Encode(resp)
		return
	}

	perfil, err := database.SearchPerfil(ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp.Ok = false
		resp.Message = "User doesn't exists."
		resp.Status = http.StatusBadRequest
		json.NewEncoder(w).Encode(resp)
		return
	}

	openFile, err := os.Open("uploads/avatar/" + perfil.Avatar)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp.Ok = false
		resp.Message = "Image hasn't been found."
		resp.Status = http.StatusBadRequest
		json.NewEncoder(w).Encode(resp)
		return
	}

	_, err = io.Copy(w, openFile)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp.Ok = false
		resp.Message = "An error ocurred during copy image."
		resp.Status = http.StatusInternalServerError
		json.NewEncoder(w).Encode(resp)
		return
	}
}
