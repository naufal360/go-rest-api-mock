package middlewares

import (
	"go-rest-api-mock/config"
	"go-rest-api-mock/helpers"
	"go-rest-api-mock/models"
	"go-rest-api-mock/repository"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		var repo = repository.NewRepo(config.GORM.DB)
		productId, err := strconv.Atoi(c.Param("productId"))
		if err != nil {
			helpers.AbortBadRequest(c, err)
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		isAdmin := userData["admin"].(bool)
		Product := models.Product{}

		Product, err = repo.GetUserIdProductById(uint(productId))
		if err != nil {
			helpers.AbortNotFound(c, err)
			return
		}

		if Product.UserID != userId && !isAdmin {
			helpers.AbortUnzuthorized(c, err)
			return
		}

		c.Next()
	}
}

func ProductUpdateDeleteAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		var repo = repository.NewRepo(config.GORM.DB)
		productId, err := strconv.Atoi(c.Param("productId"))
		if err != nil {
			helpers.AbortBadRequest(c, err)
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		isAdmin := userData["admin"].(bool)

		_, err = repo.GetUserIdProductById(uint(productId))
		if err != nil {
			helpers.AbortNotFound(c, err)
			return
		}

		if !isAdmin {
			helpers.AbortUnzuthorized(c, err)
			return
		}

		c.Next()
	}
}
