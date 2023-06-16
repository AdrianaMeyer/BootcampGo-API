package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/AdrianaMeyer/BootcampGo-API/pkg/web"
	"github.com/gin-gonic/gin"
)


func respondWithError(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, web.NewResponse(code, nil, message))
 }
 
 func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")
   
	if requiredToken == "" {
		log.Fatal("É necessário configurar a variável de ambiente TOKEN")
	}
   
	return func(c *gin.Context) {
		token := c.GetHeader("token")
	   
		if token == "" {
			respondWithError(c, http.StatusUnauthorized, "É necessário informar um token para prosseguir")
			return
		}
	   
		if token != requiredToken {
			respondWithError(c, http.StatusUnauthorized, "Token inválido")
			return
		}
	   
		c.Next()
	}
 }
 