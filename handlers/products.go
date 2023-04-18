package handlers

import (
	"go-rest-api-mock/helpers"
	"go-rest-api-mock/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (h *HttpServer) CreateProduct(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Product := models.Product{}
	userId := uint(userData["id"].(float64))

	if contentType == appJson {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userId

	data, err := h.app.CreateProduct(Product)
	if err != nil {
		helpers.CustomBadRequest(c, err)
		return
	}
	c.JSON(http.StatusCreated, data)
}

func (h *HttpServer) GetAllProduct(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Product := []models.Product{}
	userId := uint(userData["id"].(float64))
	isAdmin := userData["admin"].(bool)

	if contentType == appJson {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	data, err := h.app.GetAllProducts(userId, isAdmin)

	if err != nil {
		helpers.CustomBadRequest(c, err)
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *HttpServer) GetProductById(c *gin.Context) {
	contentType := helpers.GetContentType(c)

	Product := models.Product{}
	productId, _ := strconv.Atoi(c.Param("productId"))

	if contentType == appJson {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	data, err := h.app.GetProductById(uint(productId))

	if err != nil {
		helpers.CustomBadRequest(c, err)
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *HttpServer) UpdateProduct(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Product := models.Product{}
	productId, _ := strconv.Atoi(c.Param("productId"))
	userId := uint(userData["id"].(float64))
	isAdmin := userData["admin"].(bool)

	if contentType == appJson {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	if !isAdmin {
		Product.UserID = userId
	}
	Product.ID = uint(productId)

	data, err := h.app.UpdateProductById(uint(productId), Product)

	if err != nil {
		helpers.CustomBadRequest(c, err)
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *HttpServer) DeleteProductById(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Product := models.Product{}
	productId, _ := strconv.Atoi(c.Param("productId"))
	userId := uint(userData["id"].(float64))
	isAdmin := userData["admin"].(bool)

	if contentType == appJson {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	if !isAdmin {
		Product.UserID = userId
	}
	Product.ID = uint(productId)

	err := h.app.DeleteProductById(uint(productId))

	if err != nil {
		helpers.CustomBadRequest(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success Delete Product",
	})
}
