package tests_mocks

import (
	"encoding/json"

	"github.com/AdrianaMeyer/BootcampGo-API/internal/products"
)

type StubProductsSave struct {
}

func (s *StubProductsSave) Write(data interface{}) error {
	return nil
}
func (s *StubProductsSave) Read(data interface{}) error {
	products := []products.Product{}
	jsonData, _ := json.Marshal(products)
	return json.Unmarshal(jsonData, &data)
}