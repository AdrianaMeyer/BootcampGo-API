package main

import (
	"log"

	"github.com/AdrianaMeyer/BootcampGo-API/cmd/server/handler"
	"github.com/AdrianaMeyer/BootcampGo-API/internal/products"
	"github.com/AdrianaMeyer/BootcampGo-API/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
	  log.Fatal("Erro ao carregar o arquivo .env")
	}

	store := store.Factory("file", "../../products.json")
	if store == nil {
		log.Fatal("Não foi possível criar a Store")
	}

	repo := products.NewRepository(store)
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