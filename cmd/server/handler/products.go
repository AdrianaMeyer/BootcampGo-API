package handler

import (
	"net/http"
	"time"
	"strconv"

	"github.com/AdrianaMeyer/BootcampGo-API/internal/products"
	"github.com/gin-gonic/gin"
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
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Error": "Token Inválido",
			})
			return
		}

		product, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"Error": err.Error(),
			})
		return
		}

		ctx.JSON(http.StatusOK, product)

	}
}

func (c *Product) Save() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Error": "Token Inválido",
			})
			return
		}

		var req request
		err := ctx.Bind(&req)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"Error": err.Error(),
			})
			return	
		}
		
		date := time.Now()

		product, err := c.service.Save(req.Name, req.Color, req.Price, req.Count, req.Code, req.Published, date)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"Error": err.Error(),
			})
			return	
		}

		ctx.JSON(http.StatusCreated, product)

	}
}

func (c *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{ "Error": "Token inválido" })
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{ "Error": "ID inválido"})
			return
		}

		var req request
		err = ctx.ShouldBindJSON(&req);

		if  err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{ "Error": err.Error() })
			return
		}
		if req.Name == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{ "Error": "O campo nome do produto é obrigatório"})
			return
		}
		if req.Color == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{ "Error": "O campo cor do produto é obrigatório"})
			return
		}
		if req.Price == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{ "Error": "O campo preço do produto é obrigatório"})
			return
		}
		if req.Count == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{ "Error": "O campo quantidade é obrigatório"})
			return
		}
		if req.Code == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{ "Error": "O campo código do produto é obrigatório"})
			return
		}

		p, err := c.service.Update(int(id), req.Name, req.Color, req.Price, req.Count, req.Code, req.Published)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{ "Error": err.Error() })
			return
		}
		ctx.JSON(http.StatusOK, p)
 }}
 