package controllers

import (
	"go/auth/entities"
	"go/auth/helpers"
	"go/auth/services"
	"net/http"
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
		helpers.Response.ErrorResponse(ctx, err)
		return
	}
	if err := bc.service.CreateRecord(&book); err != nil {
		helpers.Response.ErrorResponse(ctx, err)
		return
	}
	helpers.Response.SuccessResponse(ctx, book, "Book Created")
}

func (bc *BookController) UpdateBook(ctx *gin.Context) {
	var book entities.Book
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		helpers.Response.ErrorResponse(ctx, err)
		return
	}

	err = ctx.ShouldBindJSON(&book)
	if err != nil {
		helpers.Response.ErrorResponse(ctx, err)
		return
	}

	book, err = bc.service.UpdateById(&book, id)
	if err != nil {
		helpers.Response.ErrorResponse(ctx, err)
		return
	}
	helpers.Response.SuccessResponse(ctx, book, "Book Updated")
}

func (bc *BookController) DeleteById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

	if err := bc.service.DeleteById(id); err != nil {
		helpers.Response.ErrorResponse(ctx, err)
		return
	}

	helpers.Response.SuccessResponse(ctx, nil, "Book Deleted")
}

func (bc *BookController) GetByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

	book, err := bc.service.GetByID(id)
	if err != nil {
		helpers.Response.ErrorResponse(ctx, err, http.StatusNotFound)
		return
	}
	helpers.Response.SuccessResponse(ctx, book, "Book Found")
}
