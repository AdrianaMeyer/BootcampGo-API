package mocks

import (
	"encoding/json"
	"time"

	"github.com/AdrianaMeyer/BootcampGo-API/internal/products"
)


type StubProducts struct {
}

func (s *StubProducts) Write(data interface{}) error {
	return nil
}

func (s *StubProducts) Read(data interface{}) error {
	fakeDate := time.Now()
	fakeDateFormat := fakeDate.Format("02/01/2006")

	products := []products.Product{
		{
			ID: 1,
			Name: "Estojo",
			Color: "Preto",
			Price: 30.9,
			Count: 40,
			Code: "CCC6543",
			Published: true,
			Date: fakeDateFormat,
		},
		{
			ID: 2,
			Name: "Celular",
			Color: "Azul",
			Price: 2500.9,
			Count: 40,
			Code: "CCC6542",
			Published: true,
			Date: fakeDateFormat,
		},
	}
	jsonData, _ := json.Marshal(products)
	return json.Unmarshal(jsonData, &data)
}