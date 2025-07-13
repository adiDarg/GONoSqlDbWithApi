package main

import (
	"customDatabase/go-packages/API"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/SetDoc", API.SetDoc)
	http.HandleFunc("/GetDocByID", API.GetDocByID)
	http.HandleFunc("/GetDocByName", API.GetDocByName)
	http.HandleFunc("/DeleteDocByID", API.DeleteDocByID)
	http.HandleFunc("/DeleteDocByName", API.DeleteDocByName)
	http.HandleFunc("/GenerateAPIKey", API.GenerateAPIKey)

	addr := ":4410"
	log.Printf("Starting server at %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
