package handler

import (
	"net/http"
	"time"
	"strconv"
	"fmt"

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

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products [get]
func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		product, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNoContent, web.NewResponse(http.StatusNoContent, nil, "Não há produtos cadastrados"))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, product, ""))

	}
}

// Método Save
// SaveProducts godoc
// @Summary Save products
// @Tags Products
// @Description save products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Product to save"
// @Success 200 {object} web.Response
// @Router /products [post]
func (c *Product) Save() gin.HandlerFunc {
	return func(ctx *gin.Context) {

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


// Método Update
// UpdateProducts godoc
// @Summary Update product 
// @Tags Products
// @Description Updates products based on id 
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Product to update"
// @Success 200 {object} web.Response
// @Router /products/:id [put]
func (c *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

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
 
// Método UpdateNameAndPrice
// UpdateProducts godoc
// @Summary Update Name and Price of a product 
// @Tags Products
// @Description Update Name and Price of a product based on id 
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Product name and price to update"
// @Success 200 {object} web.Response
// @Router /products/:id [patch]
func (c *Product) UpdateNameAndPrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
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
 

// Método Delete
// DeleteProducts godoc
// @Summary Delete a product
// @Tags Products
// @Description Delete a products based on id 
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Router /products/:id [delete]
func (c *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
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