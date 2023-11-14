package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Title  string
	Author string
	ISBN   string
}

func displayBook(b Book) string {
	return fmt.Sprintf("Title: %s\nAuthor: %s\nISBN: %s\n", b.Title, b.Author, b.ISBN)
}

func validateISBN(isbn string) error {
	if len(isbn) != 13 {
		return errors.New("ISBN must be 13 characters long")
	}
	return nil
}

func createBook(c *gin.Context) (Book, error) {
	title := c.PostForm("title")
	author := c.PostForm("author")
	isbn := c.PostForm("isbn")

	err := validateISBN(isbn)
	if err != nil {
		return Book{}, err
	}

	return Book{Title: title, Author: author, ISBN: isbn}, nil
}

func main() {
	router := gin.Default()

	var library []Book

	// Load HTML templates
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"library": library})
	})

	router.POST("/add", func(c *gin.Context) {
		book, err := createBook(c)
		if err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": err.Error()})
			return
		}
		library = append(library, book)
		c.Redirect(http.StatusSeeOther, "/")
	})

	router.Run(":8080")
}
