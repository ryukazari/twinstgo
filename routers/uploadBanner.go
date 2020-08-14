package routers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/ryukazari/twinstgo/database"
	"github.com/ryukazari/twinstgo/models"
)

//UploadBanner upload banner image to directory /uploads
func UploadBanner(w http.ResponseWriter, r *http.Request) {
	resp := models.Response{
		Ok:     true,
		Status: http.StatusOK,
	}
	w.Header().Set("Content-Type", "application/json")
	file, handler, err := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/banner/" + IDUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp.Ok = false
		resp.Message = "An error occurred while trying to upload the banner" + err.Error()
		resp.Status = http.StatusBadRequest
		json.NewEncoder(w).Encode(resp)
		return
	}
	_, err = io.Copy(f, file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp.Ok = false
		resp.Message = "An error occurred while trying to copy the banner" + err.Error()
		resp.Status = http.StatusBadRequest
		json.NewEncoder(w).Encode(resp)
		return
	}
	var usuario models.User
	var status bool
	usuario.Banner = IDUsuario + "." + extension
	status, err = database.ModifyRegister(usuario, IDUsuario)
	if err != nil || !status {
		w.WriteHeader(http.StatusInternalServerError)
		resp.Ok = false
		resp.Message = "An error occurred while saving banner in database" + err.Error()
		resp.Status = http.StatusInternalServerError
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(http.StatusOK)
	resp.Message = "Banner saved successfully!"
	json.NewEncoder(w).Encode(resp)
	return
}
