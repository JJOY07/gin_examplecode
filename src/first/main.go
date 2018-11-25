package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	g := router.Group("/beee")
	{
		g.GET("/health", health)
		g.POST("/signup", signup)
		g.POST("/login", login)
	}

	router.Run() // listen and serve on 0.0.0.0:8080
}

func health(cxt *gin.Context) {
	cxt.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func signup(cxt *gin.Context) {
	cxt.JSON(http.StatusCreated, gin.H{
		"message": "signed up",
	})
}

func login(cxt *gin.Context) {
	cxt.JSON(http.StatusOK, gin.H{
		"message": "logged in",
	})
}
