package controllers

import (
	"go/auth/constants"
	"go/auth/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	auth services.AuthService
}

func NewAuthController(a services.AuthService) *AuthController {
	return &AuthController{auth: a}
}

type registerDTO struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (c *AuthController) Register(ctx *gin.Context) {
	var dto registerDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := c.auth.Register(dto.Name, dto.Email, dto.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"user": user})
}

func (c *AuthController) Login(ctx *gin.Context) {
	var dto loginDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	access, refresh, err := c.auth.Login(dto.Email, dto.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"access_token": access, "refresh_token": refresh})
}

func (c *AuthController) Refresh(ctx *gin.Context) {
	var body struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	access, err := c.auth.Refresh(body.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"access_token": access})
}

func (c *AuthController) Me(ctx *gin.Context) {
	uid, exists := ctx.Get(constants.ContextUserID)
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated"})
		return
	}
	id := uid.(uint)
	user, _ := c.auth.GetByID(id)
	if user == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (c *AuthController) AdminDashboard(ctx *gin.Context) {
	uid, _ := ctx.Get(constants.ContextUserID)
	role, _ := ctx.Get(constants.ContextRole)
	ctx.JSON(http.StatusOK, gin.H{"message": "welcome admin", "user_id": strconv.FormatUint(uint64(uid.(uint)), 10), "role": role})
}
