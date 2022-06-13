package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	PORT = ":8080"
)

func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageID := vars["id"]
	fmt.Println(pageID)
	fileName := "files/" + pageID + ".html"
	http.ServeFile(w, r, fileName)
}

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/pages/{id:[0-9]+}", pageHandler)
	rtr.HandleFunc("/{id:homepage}", pageHandler)
	rtr.HandleFunc("/{id:contact}", pageHandler)
	http.Handle("/", rtr)
	http.ListenAndServe(PORT, nil)
}
