package API

import (
	"customDatabase/go-packages/Db"
	"encoding/json"
	"fmt"
	"net/http"
)

var db = Db.CreateDB()

func handleBadRequest(parameters []string, w *http.ResponseWriter) bool {
	for _, value := range parameters {
		if value == "" {
			fmt.Println("Bad Request")
			http.Error(*w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return true
		}
	}
	return false
}
func allowCrossOriginRequest(wPointer *http.ResponseWriter) {
	w := *wPointer
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
}

func CreateDoc(w http.ResponseWriter, r *http.Request) {
	allowCrossOriginRequest(&w)
	apiKey := r.URL.Query().Get("apiKey")
	name := r.URL.Query().Get("name")
	collection := r.URL.Query().Get("collection")
	if handleBadRequest([]string{apiKey, name, collection}, &w) {
		return
	}

	doc, err := db.CreateDoc(apiKey, name, collection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Doc Created")
}
func AddValueToDoc(w http.ResponseWriter, r *http.Request) {
	allowCrossOriginRequest(&w)
	apiKey := r.URL.Query().Get("apiKey")
	collection := r.URL.Query().Get("collection")
	id := r.URL.Query().Get("id")
	vName := r.URL.Query().Get("valueName")
	value := r.URL.Query().Get("value")
	if handleBadRequest([]string{apiKey, collection, id, vName, value}, &w) {
		return
	}
	err := db.AddValueToDoc(apiKey, id, collection, vName, value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println("Doc Modified - Added Value")
}
func RemoveValueFromDoc(w http.ResponseWriter, r *http.Request) {
	allowCrossOriginRequest(&w)
	apiKey := r.URL.Query().Get("apiKey")
	collection := r.URL.Query().Get("collection")
	id := r.URL.Query().Get("id")
	vName := r.URL.Query().Get("valueName")
	if handleBadRequest([]string{apiKey, collection, id, vName}, &w) {
		return
	}
	err := db.RemoveValueFromDoc(apiKey, id, collection, vName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println("Doc Modified - Removed Value")
}
func GetDocByID(w http.ResponseWriter, r *http.Request) {
	allowCrossOriginRequest(&w)
	apiKey := r.URL.Query().Get("apiKey")
	collection := r.URL.Query().Get("collection")
	id := r.URL.Query().Get("id")

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
	fmt.Println("Doc Fetched By ID")
}
func GetDocs(w http.ResponseWriter, r *http.Request) {
	allowCrossOriginRequest(&w)
	apiKey := r.URL.Query().Get("apiKey")
	collection := r.URL.Query().Get("collection")
	if handleBadRequest([]string{apiKey, collection}, &w) {
		return
	}
	docs, err := db.ReadAllDocs(apiKey, collection)
	if err != nil {
		return
	}
	err = json.NewEncoder(w).Encode(docs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("All Docs Fetched")
}
func DeleteDocByID(w http.ResponseWriter, r *http.Request) {
	allowCrossOriginRequest(&w)
	apiKey := r.URL.Query().Get("apiKey")
	collection := r.URL.Query().Get("collection")
	id := r.URL.Query().Get("id")

	if handleBadRequest([]string{apiKey, collection, id}, &w) {
		return
	}

	err := db.DeleteDocByID(apiKey, id, collection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println("Doc Deleted")
}
func GenerateAPIKey(w http.ResponseWriter, _ *http.Request) {
	allowCrossOriginRequest(&w)
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
	fmt.Println("API Key Generated")
}
