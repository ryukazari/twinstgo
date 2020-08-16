package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ryukazari/twinstgo/database"
	"github.com/ryukazari/twinstgo/models"
)

// ReadTwitRelation endopoint to get twits from my followers
func ReadTwitRelation(w http.ResponseWriter, r *http.Request) {
	resp := models.Response{
		Ok:     true,
		Status: http.StatusOK,
	}
	w.Header().Set("Content-Type", "application/json")

	if len(r.URL.Query().Get("pagina")) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		resp.Ok = false
		resp.Message = "pagina parameter is required"
		resp.Status = http.StatusBadRequest
		json.NewEncoder(w).Encode(resp)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp.Ok = false
		resp.Message = "pagina parameter must be a number greater than 0."
		resp.Status = http.StatusBadRequest
		json.NewEncoder(w).Encode(resp)
		return
	}

	respuesta, correcto := database.ReadTwitFollowers(IDUsuario, pagina)

	if !correcto {
		w.WriteHeader(http.StatusInternalServerError)
		resp.Ok = false
		resp.Message = "An error ocurred while read twits."
		resp.Status = http.StatusInternalServerError
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(http.StatusOK)
	resp.Message = "Twits has been readed successfully"
	resp.Data = respuesta
	json.NewEncoder(w).Encode(resp)
}
