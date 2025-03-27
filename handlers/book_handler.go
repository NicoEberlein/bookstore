package handlers

import (
	"fmt"
	"github.com/NicoEberlein/bookstore/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var books []models.Book = make([]models.Book, 0)
var nextId int = 0

func CreateBookHandler(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book.ID = nextId
	nextId++
	books = append(books, book)

	c.JSON(http.StatusCreated, gin.H{"id": book.ID})
}

func GetBooksHandler(c *gin.Context) {

	category := c.Query("category")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("limit", "2"))

	filteredBooks := filterBooks(books, category)

	start := (page - 1) * pageSize
	end := start + pageSize
	if start > len(filteredBooks) {
		c.JSON(http.StatusOK, []models.Book{})
		return
	}

	if end > len(filteredBooks) {
		end = len(filteredBooks)
	}

	paginatedBooks := filteredBooks[start:end]

	c.JSON(http.StatusOK, paginatedBooks)

}

func filterBooks(books []models.Book, category string) []models.Book {
	if category == "" {
		return books
	}

	filtered := []models.Book{}
	var categoryId int
	for _, c := range categories {
		if c.Name == category {
			categoryId = c.ID
		}
	}
	fmt.Println(categoryId)

	for _, book := range books {
		if book.CategoryID == categoryId {
			filtered = append(filtered, book)
		}
	}

	return filtered
}

func GetBookHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	foundBook := getBookById(id)

	if foundBook == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("book with id %d not found}", id)})
		return
	}

	c.JSON(http.StatusOK, *foundBook)

}

func DeleteBookHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	foundBook := getBookById(id)
	if foundBook == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("book with id %d not found", id)})
	}

	newBooks := make([]models.Book, len(books)-1)
	for _, book := range books {
		if book.ID != id {
			newBooks = append(newBooks, book)
		}
	}
	books = newBooks
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func UpdateBookHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedBook models.Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedBook.ID = id

	newBooks := make([]models.Book, len(books))
	for i, book := range books {
		if book.ID != id {
			newBooks[i] = book
		} else {
			newBooks[i] = updatedBook
		}
	}

	books = newBooks

	c.Status(http.StatusOK)

}

func getBookById(id int) *models.Book {
	var foundBook *models.Book = nil

	for _, book := range books {
		if book.ID == id {
			foundBook = &book
		}
	}

	return foundBook
}
