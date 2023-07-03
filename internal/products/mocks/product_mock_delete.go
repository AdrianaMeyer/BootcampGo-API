package mocks

import (
	"encoding/json"
	"time"

	"github.com/AdrianaMeyer/BootcampGo-API/internal/products"
)


type MockProductsDelete struct {
	ReadWasCalled bool
}

func (s *MockProductsDelete) Write(data interface{}) error {
	return nil
}

func (s *MockProductsDelete) Read(data interface{}) error {
	s.ReadWasCalled = true

	fakeDate := time.Now()
	fakeDateFormat := fakeDate.Format("02/01/2006")

	products := []products.Product{
		{
			ID: 1,
			Name: "Caderno",
			Color: "Colorido",
			Price: 35.99,
			Count: 240,
			Code: "CCC6590",
			Published: true,
			Date: fakeDateFormat,
		},
	}
	jsonData, _ := json.Marshal(products)
	return json.Unmarshal(jsonData, &data)
}