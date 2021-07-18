package main

import (
	"github.com/gin-gonic/gin"

	"github.com/ChernoBen/bookstore/controllers"
	"github.com/ChernoBen/bookstore/models"
)

func main() {
	r := gin.Default()
	models.ConnectDataBase()
	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.Run()
}
