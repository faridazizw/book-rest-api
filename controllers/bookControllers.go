package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var BookDatas = []Book{}

func CreateBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook.ID = fmt.Sprintf("%d", len(BookDatas)+1)
	BookDatas = append(BookDatas, newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"Book": newBook,
	})
}

func GetAllBook(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"book": BookDatas,
	})
}

func GetBook(ctx *gin.Context) {
	ID := ctx.Param("id")
	condition := false
	var bookData Book

	for i, book := range BookDatas {
		if ID == book.ID {
			condition = true
			bookData = BookDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", ID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": bookData,
	})
}

func UpdateBook(ctx *gin.Context) {
	ID := ctx.Param("id")
	condition := false
	var updatedBook Book

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range BookDatas {
		if ID == book.ID {
			condition = true
			BookDatas[i] = updatedBook
			BookDatas[i].ID = ID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error status":  "Data Not Found",
			"error_message": fmt.Sprintf("Book with id %v not found", ID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v has been successfully updated", ID),
	})
}

func DeleteBook(ctx *gin.Context) {
	ID := ctx.Param("id")
	condition := false
	var BookIndex int

	for i, book := range BookDatas {
		if ID == book.ID {
			condition = true
			BookIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", ID),
		})
		return
	}

	copy(BookDatas[BookIndex:], BookDatas[BookIndex+1:])
	BookDatas[len(BookDatas)-1] = Book{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v has been successfully deleted", ID),
	})
}
