package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ryukazari/twinstgo/database"
	"github.com/ryukazari/twinstgo/models"
)

//ReadTwits read a list of twits from database
func ReadTwits(w http.ResponseWriter, r *http.Request) {
	resp := models.Response{
		Ok:     true,
		Status: http.StatusOK,
	}
	w.Header().Set("Content-Type", "application/json")
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		resp.Ok = false
		resp.Message = "id parameter is required"
		resp.Status = http.StatusBadRequest
		json.NewEncoder(w).Encode(resp)
		return
	}

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
		resp.Message = "pagina must be a number greater than 0"
		resp.Status = http.StatusBadRequest
		json.NewEncoder(w).Encode(resp)
		return
	}

	pag := int64(pagina)
	respuesta, correcto := database.ReadTwit(ID, pag)
	if !correcto {
		w.WriteHeader(http.StatusInternalServerError)
		resp.Ok = false
		resp.Message = "Error while read twits"
		resp.Status = http.StatusInternalServerError
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp.Message = "Twits has been readed successfully"
	resp.Data = respuesta
	json.NewEncoder(w).Encode(resp)
	return
}
