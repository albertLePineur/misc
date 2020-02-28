package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Router structure
type Router struct {
	AdminUser     string
	AdminPassword string
}

// Routes Method
func (r *Router) Routes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Group("/kafka", gin.BasicAuth(gin.Accounts{r.AdminUser: r.AdminPassword}))
	router.GET("/topics", listTopics)

	return router
}

func listTopics(c *gin.Context) {
	cli := SaramaClient{
		SaramaConfig: *cfg,
	}
	c.JSON(200, gin.H{"status": http.StatusOK, "data": "prout"})
}

/*
// Router function
func Router(port int) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	authorized := router.Group("/kafka", gin.BasicAuth(gin.Accounts{"admin": "password"}))
	authorized.GET("/topics")
	{
		authorized.GET("/topics", saramaListTopics())
		//authorized.GET("/topics/:id", saramaDescribeTopic)
		//authorized.POST("/topics/:id", saramaCreateTopic)
		//authorized.DELETE("/topics/:id", saramaDeleteTopic)
	}
	router.Run("127.0.0.1:8080")
}
*/
