package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/AdrianaMeyer/BootcampGo-API/cmd/server/handler"
	"github.com/AdrianaMeyer/BootcampGo-API/internal/products"
	"github.com/AdrianaMeyer/BootcampGo-API/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type Response struct {
	Code string             `json:"code"`
	Data []products.Product `json:"data"`
}

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")
	db := store.Factory(store.FileType, "../products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)
	r := gin.Default()

	pr := r.Group("/products")
	pr.POST("/", p.Save())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.DELETE("/:id", p.Delete())
	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}
func Test_GetProduct_OK(t *testing.T) {
	r := createServer()
	request, response := createRequestTest(http.MethodGet, "/products/", "")
	r.ServeHTTP(response, request)

	var responseResult Response
	assert.Equal(t, 200, response.Code)

	err := json.Unmarshal(response.Body.Bytes(), &responseResult)

	assert.Nil(t, err)
	assert.True(t, len(responseResult.Data) > 0)
}

func Test_SaveProduct_OK(t *testing.T) {
	r := createServer()
	request, response := createRequestTest(http.MethodPost, "/products/", `{
		"name": "Teste",
		"color": "Branco",
		"price": 9.90,
		"count": 100,
		"code": "TTT4024",
		"published": true,
		"date": "03/07/2023"
	}`)

	r.ServeHTTP(response, request)
	assert.Equal(t, 201, response.Code)
}

func Test_SaveProduct_InvalidBody(t *testing.T) {
	r := createServer()
	request, response := createRequestTest(http.MethodPost, "/products/", `{
		"name": "",
		"color": "",
		"price": ,
		"count": 0,
		"code": "",
		"published": true
	}`)

	r.ServeHTTP(response, request)
	assert.Equal(t, 400, response.Code)
}

func Test_UpdateProduct_OK(t *testing.T) {
	r := createServer()
	uri := "/products/1" 
	request, response := createRequestTest(http.MethodPut, uri , `{
			"name": "Lapiseira",
			"color": "Preto",
			"price": 4.50,
			"count": 50,
			"code": "LLL6543",
			"published": true
	}`)

	r.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code)
}

func Test_DeleteProduct_OK(t *testing.T) {
	r := createServer()
	uri := "/products/8" 
	request, response := createRequestTest(http.MethodDelete, uri, "")

	r.ServeHTTP(response, request)
	assert.Equal(t, 204, response.Code)
}