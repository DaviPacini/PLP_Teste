package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Loading() {
	r := mux.NewRouter()

	r.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe(":8080", (r)))
}
