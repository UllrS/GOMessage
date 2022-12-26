package transport

import (
	objectstorage "MessageGO/ObjectStorage"
	"MessageGO/controller"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	smt, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	smt.Execute(w, nil)
}
func GetMessages(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")
	companion := r.URL.Query().Get("companion")
	startTime := r.URL.Query().Get("starttime")
	res, err := controller.GetMessages(user, companion, startTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(res)

}
func PostMessages(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("sender")
	body := r.FormValue("body")
	recipient := r.FormValue("recipient")

	fileObjStorePath := ""
	fileName := ""
	uploadfile, handle, err := r.FormFile("attachedfile")
	switch err {
	case nil:
		fileContent, err := ioutil.ReadAll(uploadfile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fileObjStorePath = fmt.Sprint(user, "_", time.Now().UnixMicro())
		fileName = handle.Filename
		err = objectstorage.SaveFile(fileContent, fileObjStorePath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.ErrMissingFile:
		fileObjStorePath = ""
		fileName = ""
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responce, err := controller.PostMessages(body, user, recipient, fileObjStorePath, fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(responce)
}
func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")
	id := r.URL.Query().Get("id")

	res, err := controller.DeleteMessage(id, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(res)
}
func DeleteChat(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")
	companion := r.URL.Query().Get("companion")

	res, err := controller.DeleteChat(user, companion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(res)
}
func GetFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get file")
	filePath := r.URL.Query().Get("path")
	fileName := r.URL.Query().Get("name")

	body, err := objectstorage.LoadFile(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	w.Write(body)
}
