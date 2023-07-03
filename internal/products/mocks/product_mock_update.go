package mocks

import (
	"encoding/json"
	"time"

	"github.com/AdrianaMeyer/BootcampGo-API/internal/products"
)


type MockProductsUpdate struct {
	ReadWasCalled bool
}

func (s *MockProductsUpdate) Write(data interface{}) error {
	return nil
}

func (s *MockProductsUpdate) Read(data interface{}) error {
	s.ReadWasCalled = true

	fakeDate := time.Now()
	fakeDateFormat := fakeDate.Format("02/01/2006")

	products := []products.Product{
		{
			ID: 8,
			Name: "NotUpdated",
			Color: "Preto",
			Price: 50000.99,
			Count: 30,
			Code: "CCC6547",
			Published: true,
			Date: fakeDateFormat,
		},
	}
	jsonData, _ := json.Marshal(products)
	return json.Unmarshal(jsonData, &data)
}