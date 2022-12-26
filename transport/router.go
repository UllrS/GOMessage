package transport

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RouterHandler(ip string) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", Index).Methods("GET")
	r.HandleFunc("/message", GetMessages).Methods("GET")
	r.HandleFunc("/message", PostMessages).Methods("POST")
	r.HandleFunc("/message/delete", DeleteMessage).Methods("GET")
	r.HandleFunc("/chat/delete", DeleteChat).Methods("GET")
	r.HandleFunc("/file", GetFile).Methods("GET")
	http.Handle("/", r)

	http.ListenAndServe(ip, nil)
	return r
}
