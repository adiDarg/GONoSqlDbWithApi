package Db

import (
	"customDatabase/go-packages/Doc"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"os"
	"strings"
	"sync"
)

const docNotFoundError string = "document not found"
const baseDir string = "Users"
const apiKeyMapDir string = "APIKeyMap"
const apiKeyMapPath = apiKeyMapDir + "/keys.json"

type SecuredDB struct {
	keyToPath map[string]string
}

var (
	instance *SecuredDB
	once     sync.Once
)

func generatePath(apiKey string, collection string, securedDB *SecuredDB) string {
	return baseDir + "/" + securedDB.keyToPath[apiKey] + "/" + collection + "/"
}

func (securedDB *SecuredDB) CreateDoc(apiKey string, name string, collection string) (Doc.Document, error) {
	doc, err := Doc.InitDoc(name)
	if err != nil {
		return Doc.Document{}, err
	}
	err = securedDB.insertDoc(apiKey, doc, collection)
	return doc, err
}
func (securedDB *SecuredDB) AddValueToDoc(apiKey string, id string, collection string, vName string, value string) error {
	doc, err := securedDB.ReadDocByID(apiKey, id, collection)
	if err != nil {
		return err
	}
	doc.AddValue(vName, value)
	return securedDB.insertDoc(apiKey, doc, collection)
}
func (securedDB *SecuredDB) RemoveValueFromDoc(apiKey string, id string, collection string, vName string) error {
	doc, err := securedDB.ReadDocByID(apiKey, id, collection)
	if err != nil {
		return err
	}
	doc.RemoveValue(vName)
	return securedDB.insertDoc(apiKey, doc, collection)
}
func (securedDB *SecuredDB) insertDoc(apiKey string, document Doc.Document, collection string) error {
	dirPath := generatePath(apiKey, collection, securedDB)
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}

	docs, err := securedDB.ReadAllDocs(apiKey, collection)
	if err != nil {
		if err.Error() != docNotFoundError {
			return err
		}
	}
	for _, dupe := range docs {
		if dupe.Name == document.Name {
			document.Id = dupe.Id
		}
	}

	jsonData, err := json.Marshal(document)
	err = os.WriteFile(dirPath+document.Id+".json", jsonData, 0644)
	return err
}
func (securedDB *SecuredDB) ReadDocByID(apiKey string, id string, collection string) (Doc.Document, error) {
	dirPath := generatePath(apiKey, collection, securedDB)
	file, err := os.Open(dirPath + id + ".json")
	if err != nil {
		return Doc.Document{}, err
	}
	defer file.Close()
	var doc Doc.Document
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&doc)
	if err != nil {
		return Doc.Document{}, err
	}
	return doc, nil
}
func (securedDB *SecuredDB) ReadAllDocs(apiKey string, collection string) ([]Doc.Document, error) {
	dirPath := generatePath(apiKey, collection, securedDB)
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return make([]Doc.Document, 0), err
	}
	documents := make([]Doc.Document, 0)
	for _, file := range files {
		nameWithoutExt := strings.TrimSuffix(file.Name(), ".json")
		doc, err := securedDB.ReadDocByID(apiKey, nameWithoutExt, collection)
		if err != nil {
			return documents, err
		}
		documents = append(documents, doc)
	}
	return documents, nil
}
func (securedDB *SecuredDB) DeleteDocByID(apiKey string, id string, collection string) error {
	dirPath := generatePath(apiKey, collection, securedDB)
	err := os.Remove(dirPath + id + ".json")
	return err
}
func (securedDB *SecuredDB) GenerateAPIKey() (string, error) {
	userKey := uuid.NewString()
	filePath := uuid.NewString()
	securedDB.keyToPath[userKey] = filePath
	err := securedDB.saveAPIKeyMap()
	return userKey, err
}
func (securedDB *SecuredDB) saveAPIKeyMap() error {
	err := os.MkdirAll(apiKeyMapDir, os.ModePerm)
	if err != nil {
		return err
	}
	jsonData, err := json.Marshal(securedDB.keyToPath)
	if err != nil {
		return err
	}
	err = os.WriteFile(apiKeyMapPath, jsonData, 0644)
	return err
}
func loadAPIKeyMap() (map[string]string, error) {
	data, err := os.ReadFile(apiKeyMapPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return make(map[string]string), nil
		}
		return nil, err
	}
	var retObj map[string]string
	err = json.Unmarshal(data, &retObj)
	return retObj, err
}

func CreateDB() *SecuredDB {
	once.Do(func() {
		keyToPath, err := loadAPIKeyMap()
		if err != nil {
			keyToPath = make(map[string]string)
		}
		instance = &SecuredDB{keyToPath: keyToPath}
	})
	return instance
}
