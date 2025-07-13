package API

import (
	"customDatabase/go-packages/Db"
	"customDatabase/go-packages/Doc"
	"encoding/json"
	"net/http"
)

var db = Db.CreateDB()

func handleBadRequest(parameters []string, w *http.ResponseWriter) bool {
	for _, value := range parameters {
		if value == "" {
			http.Error(*w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return true
		}
	}
	return false
}

func SetDoc(w http.ResponseWriter, r *http.Request) {
	apiKey := r.URL.Query().Get("apiKey")
	collection := r.URL.Query().Get("collection")

	if handleBadRequest([]string{apiKey, collection}, &w) {
		return
	}

	var document Doc.Document
	err := json.NewDecoder(r.Body).Decode(&document)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = db.InsertDoc(apiKey, document, collection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func GetDocByID(w http.ResponseWriter, r *http.Request) {
	apiKey := r.URL.Query().Get("apiKey")
	collection := r.URL.Query().Get("collection")
	id := r.URL.Query().Get("documentId")

	if handleBadRequest([]string{apiKey, collection, id}, &w) {
		return
	}

	doc, err := db.ReadDocByID(apiKey, id, collection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func GetDocByName(w http.ResponseWriter, r *http.Request) {
	apiKey := r.URL.Query().Get("apiKey")
	collection := r.URL.Query().Get("collection")
	name := r.URL.Query().Get("name")

	if handleBadRequest([]string{apiKey, collection, name}, &w) {
		return
	}

	doc, err := db.ReadDocByName(apiKey, name, collection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func DeleteDocByID(w http.ResponseWriter, r *http.Request) {
	apiKey := r.URL.Query().Get("apiKey")
	collection := r.URL.Query().Get("collection")
	id := r.URL.Query().Get("documentId")

	if handleBadRequest([]string{apiKey, collection, id}, &w) {
		return
	}

	err := db.DeleteDocByID(apiKey, id, collection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func DeleteDocByName(w http.ResponseWriter, r *http.Request) {
	apiKey := r.URL.Query().Get("apiKey")
	collection := r.URL.Query().Get("collection")
	name := r.URL.Query().Get("name")

	if handleBadRequest([]string{apiKey, collection, name}, &w) {
		return
	}

	err := db.DeleteDocByID(apiKey, name, collection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func GenerateAPIKey(w http.ResponseWriter, r *http.Request) {
	key, err := db.GenerateAPIKey()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(key))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
