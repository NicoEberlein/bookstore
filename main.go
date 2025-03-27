package main

import (
	"github.com/NicoEberlein/bookstore/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.GET("/books", handlers.GetBooksHandler)
	router.GET("/books/:id", handlers.GetBookHandler)
	router.GET("/categories", handlers.GetCategoriesHandler)
	router.GET("/authors", handlers.GetAuthorsHandler)

	router.POST("/books", handlers.CreateBookHandler)
	router.POST("/categories", handlers.CreateCategoryHandler)
	router.POST("/authors", handlers.CreateAuthorHandler)

	router.DELETE("/books/:id", handlers.DeleteBookHandler)

	router.PUT("/books/:id", handlers.UpdateBookHandler)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
