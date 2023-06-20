package tests

import (
	"testing"
	"time"

	"github.com/AdrianaMeyer/BootcampGo-API/internal/products"
	tests_mocks "github.com/AdrianaMeyer/BootcampGo-API/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	fakeDate := time.Now()
	fakeDateFormat := fakeDate.Format("02/01/2006")

	expectedResult := []products.Product{
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

	MyStub := tests_mocks.StubProducts{}
	MyRepoMock := products.NewRepository(&MyStub)
	MyService := products.NewService(MyRepoMock)
	result, err := MyService.GetAll()

	assert.Nil(t, err)
	assert.NoError(t, err)
	assert.Equal(t, result, expectedResult)
}

func TestGetAllError(t *testing.T) {
	MyStub := tests_mocks.StubProductsError{}
	MyRepoMock := products.NewRepository(&MyStub)
	MyService := products.NewService(MyRepoMock)
	result, err := MyService.GetAll()

	expectedError := "Erro ao ler os dados"

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), expectedError)
}

func TestSave(t *testing.T) {
	fakeDate := time.Now()
	fakeDateFormat := fakeDate.Format("02/01/2006")

	testProduct := products.Product{
		ID: 1,
		Name: "Estojo",
		Color: "Preto",
		Price: 30.9,
		Count: 40,
		Code: "CCC6543",
		Published: true,
		Date: fakeDateFormat,
	}

	MyStub := tests_mocks.StubProductsSave{}
	MyRepoMock := products.NewRepository(&MyStub)
	MyService := products.NewService(MyRepoMock)
	result, err := MyService.Save(
		testProduct.Name,
		testProduct.Color,
		testProduct.Price,
		testProduct.Count,
		testProduct.Code,
		testProduct.Published,
		testProduct.Date,
	)

	assert.Nil(t, err)
	assert.Equal(t, result.ID, 1)
}

func TestSaveError(t *testing.T) {
	testProduct := products.Product{}
	expectedErrorMessage := "JSON unexpected character"

	MyStub := tests_mocks.StubProductsSaveError{}
	MyRepoMock := products.NewRepository(&MyStub)
	MyService := products.NewService(MyRepoMock)
	_, err := MyService.Save(
		testProduct.Name,
		testProduct.Color,
		testProduct.Price,
		testProduct.Count,
		testProduct.Code,
		testProduct.Published,
		testProduct.Date,
	)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), expectedErrorMessage)
}