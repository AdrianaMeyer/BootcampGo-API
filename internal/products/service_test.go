package products_test

import (
	"testing"
	"time"
	"errors"

	"github.com/AdrianaMeyer/BootcampGo-API/internal/products"
	"github.com/AdrianaMeyer/BootcampGo-API/internal/products/mocks"
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

	MyStub := mocks.StubProducts{}
	MyRepoMock := products.NewRepository(&MyStub)
	MyService := products.NewService(MyRepoMock)
	result, err := MyService.GetAll()

	assert.Nil(t, err)
	assert.NoError(t, err)
	assert.Equal(t, result, expectedResult)
}

func TestGetAllError(t *testing.T) {
	MyStub := mocks.StubProductsError{}
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

	MyStub := mocks.StubProductsSave{}
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

	MyStub := mocks.StubProductsSaveError{}
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

func TestUpdateNameAndPrice(t *testing.T) {

	updatedProduct := products.Product{
		ID: 1,
		Name: "After Update",
		Price: 30.9,
	}

	MyMock := mocks.MockProductsUpdateNamePrice{}
	MyRepoMock := products.NewRepository(&MyMock)
	MyService := products.NewService(MyRepoMock)
	result, err := MyService.UpdateNameAndPrice(
		updatedProduct.ID,
		updatedProduct.Name,
		updatedProduct.Price,
	)

	assert.Nil(t, err)
	assert.True(t, result.Name == "After Update")
	assert.True(t, result.Price == 30.9)
	assert.True(t, MyMock.ReadWasCalled)
}

func TestUpdateNameAndPriceError(t *testing.T) {

	updatedProduct := products.Product{
		ID: 99,
		Name: "After Update",
		Price: 30.9,
	}

	MyMock := mocks.MockProductsUpdateNamePrice{}
	MyRepoMock := products.NewRepository(&MyMock)
	MyService := products.NewService(MyRepoMock)
	_, err := MyService.UpdateNameAndPrice(
		updatedProduct.ID,
		updatedProduct.Name,
		updatedProduct.Price,
	)

	expectedError := errors.New("produto não encontrado")

	assert.Equal(t, expectedError, err)
	assert.Error(t, expectedError)
	assert.True(t, MyMock.ReadWasCalled)
}

func TestUpdate(t *testing.T) {

	updatedProduct := products.Product{
		ID: 8,
		Name: "Updated",
		Color: "Amarelo",
		Price: 34.80,
		Count: 45,
		Code: "CCC6548",
		Published: false,
	}

	MyMock := mocks.MockProductsUpdate{}
	MyRepoMock := products.NewRepository(&MyMock)
	MyService := products.NewService(MyRepoMock)
	result, err := MyService.Update(
		updatedProduct.ID,
		updatedProduct.Name,
		updatedProduct.Color,
		updatedProduct.Price,
		updatedProduct.Count,
		updatedProduct.Code,
		updatedProduct.Published,
	)

	assert.Nil(t, err)
	assert.True(t, result.ID == updatedProduct.ID)
	assert.True(t, result.Name == updatedProduct.Name)
	assert.True(t, result.Color == updatedProduct.Color)
	assert.True(t, result.Price == updatedProduct.Price)
	assert.True(t, result.Count == updatedProduct.Count)
	assert.True(t, result.Code == updatedProduct.Code)
	assert.True(t, result.Published == updatedProduct.Published)
	assert.True(t, MyMock.ReadWasCalled)
}

func TestUpdateError(t *testing.T) {

	updatedProduct := products.Product{
		ID: 99,
		Name: "Updated",
		Color: "Amarelo",
		Price: 34.80,
		Count: 45,
		Code: "CCC6548",
		Published: false,
	}

	MyMock := mocks.MockProductsUpdate{}
	MyRepoMock := products.NewRepository(&MyMock)
	MyService := products.NewService(MyRepoMock)
	_, err := MyService.Update(
		updatedProduct.ID,
		updatedProduct.Name,
		updatedProduct.Color,
		updatedProduct.Price,
		updatedProduct.Count,
		updatedProduct.Code,
		updatedProduct.Published,
	)

	expectedError := errors.New("produto não encontrado")

	assert.Equal(t, expectedError, err)
	assert.Error(t, expectedError)
	assert.True(t, MyMock.ReadWasCalled)
}

func TestDelete(t *testing.T) {

	IDExists := 1

	MyMock := mocks.MockProductsDelete{}
	MyRepoMock := products.NewRepository(&MyMock)
	MyService := products.NewService(MyRepoMock)
	result := MyService.Delete(IDExists)

	assert.Nil(t, result)
	assert.True(t, MyMock.ReadWasCalled)
}

func TestDeleteError(t *testing.T) {

	IDNotExist := 99

	MyMock := mocks.MockProductsDelete{}
	MyRepoMock := products.NewRepository(&MyMock)
	MyService := products.NewService(MyRepoMock)
	result := MyService.Delete(IDNotExist)

	expectedError := errors.New("produto não encontrado")

	assert.Equal(t, expectedError, result)
	assert.True(t, MyMock.ReadWasCalled)
}