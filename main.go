package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const httpPort = 1337

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", RootEndpoint).Methods("GET")

	log.Print(fmt.Sprintf("Starting server at %d", httpPort))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", httpPort), router))
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
