package handler

import (
	"net/http"
	"time"
	"strconv"
	"fmt"
	"os"

	"github.com/AdrianaMeyer/BootcampGo-API/internal/products"
	"github.com/gin-gonic/gin"
	"github.com/AdrianaMeyer/BootcampGo-API/pkg/web"
)

type request struct {
	Name 		string 		`json:"name"`
	Color 		string 		`json:"color"`
	Price 		float64 	`json:"price"`
	Count 		int 		`json:"count"`
	Code 		string 		`json:"code"`
	Published 	bool 		`json:"published"`
	Date 		time.Time 	`json:"date"`
}

type Product struct {
	service products.IService

}

func NewProduct(p products.IService) *Product {
	return &Product {
		service: p,
	}
}

func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN"){
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token Inválido"))
			return
		}

		product, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNoContent, web.NewResponse(http.StatusNoContent, nil, "Não há produtos cadastrados"))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, product, ""))

	}
}

func (c *Product) Save() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token Inválido"))
			return
		}

		var req request
		err := ctx.Bind(&req)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, err.Error()))
			return	
		}

		invalidFields := validateFields(req)
		if invalidFields != nil {
			message := fmt.Sprintf("Existem campos obrigatórios que devem ser preenchidos: %v", invalidFields)
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, message))
			return
		}
		
		date := time.Now()

		product, err := c.service.Save(req.Name, req.Color, req.Price, req.Count, req.Code, req.Published, date)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
			return	
		}

		ctx.JSON(http.StatusCreated, product)

	}
}

func (c *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token Inválido"))
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 0)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "ID Inválido"))
			return
		}

		var req request
		err = ctx.ShouldBindJSON(&req);

		if  err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, err.Error()))
			return
		}
		
		invalidFields := validateFields(req)
		if invalidFields != nil {
			message := fmt.Sprintf("Existem campos obrigatórios que devem ser preenchidos: %v", invalidFields)
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, message))
			return
		}

		p, err := c.service.Update(int(id), req.Name, req.Color, req.Price, req.Count, req.Code, req.Published)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
 }}
 

 func (c *Product) UpdateNameAndPrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token Inválido"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"),10, 0)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "ID Inválido"))
			return
		}

		var req request
		err = ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, err.Error()))
			return
		}

		if req.Name == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "O nome do produto é obrigatório"))
			return
		}

		if req.Price == 0 {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "O preço do produto é obrigatório"))
			return
		}

		p, err := c.service.UpdateNameAndPrice(int(id), req.Name, req.Price)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, p)
 }}
 
 func (c *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN"){
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token Inválido"))
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"),10, 0)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "ID Inválido"))
			return
		}

		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusNoContent, gin.H{ "Sucesso!": fmt.Sprintf("O produto %d foi removido", id) })
	}
 }
 
 func validateFields(req request) []string {

	emptyfields := []string{}

	if req.Name == "" {
		emptyfields = append(emptyfields, "Name")
	}

	if req.Color == "" {
		emptyfields = append(emptyfields, "Color")
	}

	if req.Price == 0 {
		emptyfields = append(emptyfields, "Price")
	}

	if req.Count == 0 {
		emptyfields = append(emptyfields, "Count")
	}

	if req.Code == "" {
		emptyfields = append(emptyfields, "Code")
	}

	if len(emptyfields) != 0 {
		return emptyfields
	} else {
		return nil
	}

 }