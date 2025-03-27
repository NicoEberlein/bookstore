package handlers

import (
	"github.com/NicoEberlein/bookstore/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var categories = make([]models.Category, 0)
var nextCategoryId = 0

func CreateCategoryHandler(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category.ID = nextCategoryId
	nextCategoryId++

	categories = append(categories, category)

	c.JSON(http.StatusCreated, gin.H{"id": category.ID})
}

func GetCategoriesHandler(c *gin.Context) {
	if len(categories) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no categories found"})
		return
	}

	c.JSON(http.StatusOK, categories)
	return
}
