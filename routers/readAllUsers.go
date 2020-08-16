package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ryukazari/twinstgo/database"
	"github.com/ryukazari/twinstgo/models"
)

//ReadAllUsers endpoint to read all of users from database
func ReadAllUsers(w http.ResponseWriter, r *http.Request) {
	resp := models.Response{
		Ok:     true,
		Status: http.StatusOK,
	}
	w.Header().Set("Content-Type", "application/json")

	typeUser := r.URL.Query().Get("type")
	pagina := r.URL.Query().Get("pagina")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(pagina)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp.Ok = false
		resp.Message = "Must send pagina parameter as a number greater than 0."
		resp.Status = http.StatusBadRequest
		json.NewEncoder(w).Encode(resp)
		return
	}

	pag := int64(pagTemp)
	result, status := database.ReadAllUsers(IDUsuario, pag, search, typeUser)
	if !status {
		w.WriteHeader(http.StatusInternalServerError)
		resp.Ok = false
		resp.Message = "An error ocurred while reading users."
		resp.Status = http.StatusInternalServerError
		json.NewEncoder(w).Encode(resp)
		return
	}
	w.WriteHeader(http.StatusOK)
	resp.Message = "Users were obtained successfully!"
	resp.Data = result
	json.NewEncoder(w).Encode(resp)
	return
}
