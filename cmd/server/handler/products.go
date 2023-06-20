package handler

import (
	"net/http"
	"time"
	"strconv"
	"fmt"

	"github.com/AdrianaMeyer/BootcampGo-API/internal/products"
	"github.com/AdrianaMeyer/BootcampGo-API/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/AdrianaMeyer/BootcampGo-API/pkg/web"
)

type Product struct {
	service products.IService

}

func NewProduct(p products.IService) *Product {
	return &Product {
		service: p,
	}
}

// GetAll godoc
// @Summary List all products
// @Description getAll products
// @Tags Products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Failure 204 {object} web.Response "Não há produtos cadastrados"
// @Router /products/ [get]
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

// Save godoc
// @Summary Save new products
// @Description Create a new product based on the provided JSON
// @Tags Products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body domain.Request true "Product to be saves"
// @Success 201 {object} web.Response "Created product"
// @Failure 400 {object} web.Response "Missing fields error"
// @Failure 422 {object} web.Response "Json Parse error"
// @Failure 500 {object} web.Response "Internal Server error"
// @Router /products/ [post]
func (c *Product) Save() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req domain.Request
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
		
		Date := time.Now()
		DateFormat := Date.Format("02/01/2006")

		product, err := c.service.Save(req.Name, req.Color, req.Price, req.Count, req.Code, req.Published, DateFormat)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
			return	
		}

		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, product, ""))

	}
}

// Update godoc
// @Summary Update a product based on ID
// @Description Update a specific product based on the provided JSON
// @Tags Products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body domain.Request true "Product to be updated"
// @Param id path int true "Product ID"
// @Success 200 {object} web.Response "Product Updated"
// @Failure 400 {object} web.Response  "ID validation error or missing fields"
// @Failure 404 {object} web.Response  "Product ID not found"
// @Failure 422 {object} web.Response  "Json Parse error"
// @Router /products/{id} [put]
func (c *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 0)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "ID Inválido"))
			return
		}

		var req domain.Request
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
 
// UpdateNameAndPrice godoc
// @Summary Update a product`s name and price based on ID
// @Description Update a specific product based on the provided JSON
// @Tags Products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body domain.RequestUpdateNameAndPrice true "Product to be updated"
// @Param id path int true "Product ID"
// @Success 200 {object} web.Response  "Product Updated"
// @Failure 400 {object} web.Response  "ID validation error or missing fields"
// @Failure 404 {object} web.Response  "Product ID not found"
// @Failure 422 {object} web.Response  "Json Parse error"
// @Router /products/{id} [patch]
func (c *Product) UpdateNameAndPrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"),10, 0)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "ID Inválido"))
			return
		}

		var req domain.RequestUpdateNameAndPrice
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
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
}}
 

// Delete godoc
// @Summary Delete a product based on OD
// @Description Delete a specific product based on ID
// @Tags Products
// @Param token header string true "token"
// @Param id path int true "Product ID"
// @Success 204 {object} web.Response "No content"
// @Failure 400 {object} web.Response "ID validation error"
// @Failure 404 {object} web.Response "Product not found"
// @Router /products/{id} [delete]
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
		ctx.JSON(http.StatusNoContent, web.NewResponse(http.StatusNoContent, "Produto removido com sucesso", ""))
	}
}
 
func validateFields(req domain.Request) []string {

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