package main

import (
	"log"
	"os"

	"github.com/AdrianaMeyer/BootcampGo-API/cmd/server/handler"
	"github.com/AdrianaMeyer/BootcampGo-API/internal/products"
	"github.com/AdrianaMeyer/BootcampGo-API/pkg/store"
	"github.com/AdrianaMeyer/BootcampGo-API/docs"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


// @title Bootcamp Go - API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url adriana.meyer@mercadolivre.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
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
	{
		pr.Use(handler.TokenAuthMiddleware())
		
		pr.POST("/", product.Save())
		pr.GET("/", product.GetAll())
		pr.PUT("/:id", product.Update())
		pr.PATCH("/:id", product.UpdateNameAndPrice())
		pr.DELETE("/:id", product.Delete())
	}

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run()
}