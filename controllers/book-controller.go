package controllers

import (
	"go/auth/entities"
	"go/auth/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	service services.BookService
}

func NewBookController(service services.BookService) *BookController {
	return &BookController{
		service: service,
	}
}

func (bc *BookController) CreateBook(ctx *gin.Context) {
	var book entities.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := bc.service.CreateRecord(&book); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, gin.H{"id": book.ID})
}

func (bc *BookController) UpdateBook(ctx *gin.Context) {
	var book entities.Book
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = ctx.ShouldBindJSON(&book)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = bc.service.UpdateById(&book, id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, "Book Updated")
}

func (bc *BookController) DeleteById(ctx *gin.Context) {

	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

	if err := bc.service.DeleteById(id); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"id": id})
}

func (bc *BookController) GetByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

	book, err := bc.service.GetByID(id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, book)
}
