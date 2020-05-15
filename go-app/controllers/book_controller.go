package controllers

import (
	"go-app/domain/book"
	"go-app/services/bookservice"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BookOutput represents HTTP Response Body structure
type BookOutput struct {
	Name  string `json:"name"`
	Pages uint   `json:"pages"`
}

// BookInput represents postBook body format
type BookInput struct {
	Name  string `json:"name"`
	Pages uint   `json:"pages"`
}

// BookController interface
type BookController interface {
	PostBook(*gin.Context)
}

type bookController struct {
	bs bookservice.BookService
}

// NewBookController instantiates Book Controller
func NewBookController(bs bookservice.BookService) BookController {
	return &bookController{bs: bs}
}

func (ctl *bookController) PostBook(c *gin.Context) {
	// Read user input
	var bookInput BookInput
	if err := c.ShouldBindJSON(&bookInput); err != nil {
		HTTPRes(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	b := ctl.inputToBook(bookInput)

	// Create book
	// If an Error Occurs while creating return the error
	if _, err := ctl.bs.CreateBook(&b); err != nil {
		HTTPRes(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// If book is successfully created return a structured Response
	bookOutput := ctl.mapToBookOutput(&b)
	HTTPRes(c, http.StatusOK, "Book Published", bookOutput)
}

// Private Methods
func (ctl *bookController) inputToBook(input BookInput) book.Book {
	return book.Book{
		Name:  input.Name,
		Pages: input.Pages,
	}
}
func (ctl *bookController) mapToBookOutput(b *book.Book) *BookOutput {
	return &BookOutput{
		Name:  b.Name,
		Pages: b.Pages,
	}
}
