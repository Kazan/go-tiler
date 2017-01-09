package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", RootEndpoint).Methods("GET")

	log.Fatal(http.ListenAndServe(":1337", router))
}

// RootEndpoint ...
//
func RootEndpoint(w http.ResponseWriter, req *http.Request) {
	var payload [3]string

	payload[0] = "1"
	payload[1] = "2"
	payload[2] = "3"

	json.NewEncoder(w).Encode(payload)
}
