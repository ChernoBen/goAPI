package main

import (
	"github.com/gin-gonic/gin"

	"github.com/ChernoBen/bookstore/controllers"
	"github.com/ChernoBen/bookstore/models"
)

func main() {
	r := gin.Default()
	models.ConnectDataBase()

	r.GET("/books/:id", controllers.FindBook)
	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.Delete)

	r.Run()
}
