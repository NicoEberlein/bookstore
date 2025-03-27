package handlers

import (
	"github.com/NicoEberlein/bookstore/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var authors = make([]models.Author, 0)
var nextAuthorId = 0

func CreateAuthorHandler(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author.ID = nextAuthorId
	nextAuthorId++

	authors = append(authors, author)

	c.JSON(http.StatusCreated, gin.H{"id": author.ID})
}

func GetAuthorsHandler(c *gin.Context) {
	if len(authors) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no authors found"})
		return
	}

	c.JSON(http.StatusOK, authors)
	return
}
