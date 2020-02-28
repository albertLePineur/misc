package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	routeur := gin.New()
	routeur.Use(gin.Logger())
	routeur.Use(gin.Recovery())
	authorized := routeur.Group("/", gin.BasicAuth(gin.Accounts{"admin": "admin"}))
	authorized.GET("/coucou", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"first name": "thibaud",
			"last name":  "coueffe",
		})

	})

	routeur.Run(":8080")
}
