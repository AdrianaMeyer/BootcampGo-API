package main

import (
	"github.com/AdrianaMeyer/BootcampGo-API/cmd/server/handler"
	"github.com/AdrianaMeyer/BootcampGo-API/internal/products"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
	  log.Fatal("Erro ao carregar o arquivo .env")
	}

	repo := products.NewRepository()
	service := products.NewService(repo)
	product := handler.NewProduct(service)

	router := gin.Default()
	pr := router.Group("/products")
	pr.POST("/", product.Save())
	pr.GET("/", product.GetAll())
	pr.PUT("/:id", product.Update())
	pr.PATCH("/:id", product.UpdateNameAndPrice())
	pr.DELETE("/:id", product.Delete())

	router.Run()
}