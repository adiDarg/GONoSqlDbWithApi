package Doc

import (
	"errors"
	"github.com/google/uuid"
)

type Document struct {
	Id     string
	Name   string
	Values map[string]string
}

func (doc *Document) GetValue() (string, error) {
	value, exists := doc.Values[doc.Name]
	if !exists {
		return "", errors.New("document does not exist")
	}
	return value, nil
}
func (doc *Document) AddValue(name string, value string) {
	doc.Values[name] = value
}
func (doc *Document) GetKeys() []string {
	keys := make([]string, len(doc.Values))
	i := 0
	for k := range doc.Values {
		keys[i] = k
		i++
	}
	return keys
}
func (doc *Document) createGUID() error {
	newUUID, err := uuid.NewUUID()
	if err == nil {
		doc.Id = newUUID.String()
	}
	return err
}
func InitDoc(name string) (Document, error) {
	doc := Document{"", name, map[string]string{}}
	err := doc.createGUID()
	return doc, err
}
