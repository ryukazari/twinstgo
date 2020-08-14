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

//UploadAvatar upload avatar image to directory /uploads
func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	resp := models.Response{
		Ok:     true,
		Status: http.StatusOK,
	}
	w.Header().Set("Content-Type", "application/json")
	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/avatar/" + IDUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp.Ok = false
		resp.Message = "An error occurred while trying to upload the avatar" + err.Error()
		resp.Status = http.StatusBadRequest
		json.NewEncoder(w).Encode(resp)
		return
	}
	_, err = io.Copy(f, file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp.Ok = false
		resp.Message = "An error occurred while trying to copy the avatar" + err.Error()
		resp.Status = http.StatusBadRequest
		json.NewEncoder(w).Encode(resp)
		return
	}
	var usuario models.User
	var status bool
	usuario.Avatar = IDUsuario + "." + extension
	status, err = database.ModifyRegister(usuario, IDUsuario)
	if err != nil || !status {
		w.WriteHeader(http.StatusInternalServerError)
		resp.Ok = false
		resp.Message = "An error occurred while saving avatar in database" + err.Error()
		resp.Status = http.StatusInternalServerError
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(http.StatusOK)
	resp.Message = "Avatar saved successfully!"
	json.NewEncoder(w).Encode(resp)
	return
}
