package handlers

import (
	"go-rest-api-mock/helpers"
	"go-rest-api-mock/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CreateProduct godoc
// @Summary Post a new product data
// @Description Post details of product corresponding to the input Id
// @Tags products
// @Accept json
// @Produce json
// @Param Authorization header string true "With the bearer started"
// @Param models.Product body models.Product true "create products"
// @Success 201 {object} models.Product
// @Router /products [post]
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

// GetAllProduct godoc
// @Summary Get all products data
// @Description Get details of all products
// @Tags products
// @Accept json
// @Produce json
// @Param Authorization header string true "With the bearer started"
// @Success 200 {object} models.Product
// @Router /products [get]
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

// GetProductById godoc
// @Summary Get product data for a given Id
// @Description Get details of product corresponding to the input Id
// @Tags products
// @Accept json
// @Produce json
// @Param Authorization header string true "With the bearer started"
// @Success 200 {object} models.Product
// @Router /products/{productId} [get]
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

// UpdateProduct godoc
// @Summary Update products data by admin
// @Description Update details of product corresponding to the input Id
// @Tags products
// @Accept json
// @Produce json
// @Param Authorization header string true "With the bearer started"
// @Param models.Product body models.Product true "update product"
// @Success 200 {object} models.Product
// @Router /products/{productId} [put]
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

// DeleteProduct godoc
// @Summary Update products data by admin
// @Description Delete details of product corresponding to the input Id
// @Tags products
// @Accept json
// @Produce json
// @Param Authorization header string true "With the bearer started"
// @Success 200 {object} object{message=string}
// @Router /products/{productId} [delete]
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
