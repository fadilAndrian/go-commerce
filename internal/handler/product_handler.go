package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/fadilAndrian/go-commerce/internal/helper"
	"github.com/fadilAndrian/go-commerce/internal/product"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *product.ProductService
}

func NewProductHandler(s *product.ProductService) *ProductHandler {
	return &ProductHandler{s}
}

func (handler *ProductHandler) List(c *gin.Context) {
	products, err := handler.service.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Succeed",
		"data":    products,
	})
}

func (handler *ProductHandler) Create(c *gin.Context) {
	var request *product.ProductRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	if errors := helper.ValidateRequest(request); errors != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": errors,
		})
		return
	}

	if err := handler.service.Create(c.Request.Context(), request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Product Created",
	})
}

func (handler *ProductHandler) Show(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Invalid product id",
		})
		return
	}

	product, err := handler.service.Show(c.Request.Context(), int64(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{
				"err": "Data not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Succeed",
		"data":    product,
	})
}

func (handler *ProductHandler) Update(c *gin.Context) {
	var request *product.ProductRequest

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Invalid product id",
		})
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	if errors := helper.ValidateRequest(request); errors != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": errors,
		})
		return
	}

	if err := handler.service.Update(c.Request.Context(), int64(id), request); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{
				"err": "Data not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product Updated",
	})
}

func (handler *ProductHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Invalid product id",
		})
		return
	}

	if err := handler.service.Delete(c.Request.Context(), int64(id)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{
				"err": "Data not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product Deleted",
	})
}
