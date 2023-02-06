package main

import (
	"log"
    "net/http"
	"github.com/gorilla/mux"
	"path"
    "path/filepath"
)

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/create", create).Methods("POST")
    router.HandleFunc("/read", read).Methods("GET")
    router.HandleFunc("/update", update).Methods("PUT")
    router.HandleFunc("/delete", delete_).Methods("DELETE")

    err := http.ListenAndServe(":8080", router)
    if err != nil {
        log.Fatalln("There's an error with the server," err)
    }
}