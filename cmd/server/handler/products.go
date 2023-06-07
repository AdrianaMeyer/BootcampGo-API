package handler

import (
	"net/http"
	"time"

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