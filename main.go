package main

import (
	"customDatabase/go-packages/API"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/CreateDoc", API.CreateDoc)
	http.HandleFunc("/AddValueToDoc", API.AddValueToDoc)
	http.HandleFunc("/RemoveValueFromDoc", API.RemoveValueFromDoc)
	http.HandleFunc("/GetDocs", API.GetDocs)
	http.HandleFunc("/GetDocByID", API.GetDocByID)
	http.HandleFunc("/DeleteDocByID", API.DeleteDocByID)
	http.HandleFunc("/GenerateAPIKey", API.GenerateAPIKey)
	addr := ":4410"
	log.Printf("Starting server at %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
