package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/fadilAndrian/go-commerce/internal/user"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *user.UserService
}

func NewUserHandler(service *user.UserService) *UserHandler {
	return &UserHandler{service}
}

func (handler *UserHandler) Register(c *gin.Context) {
	var request *user.RegisterUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Failed to parse request, err: " + err.Error(),
		})
		return
	}

	if err := handler.service.Register(c.Request.Context(), request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "Failed to register, err: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Register Succeed",
	})
}

func (handler *UserHandler) Login(c *gin.Context) {
	var request *user.LoginUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Failed to parse request, err: " + err.Error(),
		})
		return
	}

	token, err := handler.service.Login(c.Request.Context(), request)
	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": "Invalid Email or Password",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "Failed to login, err: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Succeed",
		"data": gin.H{
			"token": "Bearer " + token,
		},
	})
}

func (handler *UserHandler) Me(c *gin.Context) {
	userId, existed := c.Get("userId")
	if !existed {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": "unauthorized",
		})
		return
	}

	userIdInt64, ok := userId.(int64)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Invalid user id",
		})
		return
	}

	user, err := handler.service.AuthProfile(c.Request.Context(), userIdInt64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Succeed",
		"data":    user,
	})
}
