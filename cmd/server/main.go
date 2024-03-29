package main

import (
	"log"
	"os"

	"github.com/AdrianaMeyer/BootcampGo-API/cmd/server/handler"
	"github.com/AdrianaMeyer/BootcampGo-API/config"
	"github.com/AdrianaMeyer/BootcampGo-API/docs"
	"github.com/AdrianaMeyer/BootcampGo-API/internal/products"
	"github.com/AdrianaMeyer/BootcampGo-API/pkg/store"
	"github.com/gin-gonic/gin"
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
	config.InitConfig()

	store := store.Factory("file", "../../products.json")
	if store == nil {
		log.Fatal("Não foi possível criar a Store")
	}

	repo := products.NewRepository(store)
	service := products.NewService(repo)
	product := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")
	{
		pr.Use(handler.TokenAuthMiddleware())
		
		pr.POST("/", product.Save())
		pr.GET("/", product.GetAll())
		pr.PUT("/:id", product.Update())
		pr.PATCH("/:id", product.UpdateNameAndPrice())
		pr.DELETE("/:id", product.Delete())
	}

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
	
}