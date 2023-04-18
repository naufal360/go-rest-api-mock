package handlers

import (
	"go-rest-api-mock/helpers"
	"go-rest-api-mock/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJson = "application/json"
)

func (h *HttpServer) UserRegister(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	User := models.User{}

	if contentType == appJson {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	data, err := h.app.RegisterUsers(User)
	if err != nil {
		helpers.CustomBadRequest(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        data.ID,
		"email":     data.Email,
		"full_name": data.Fullname,
	})
}

func (h *HttpServer) UserLogin(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	User := models.User{}

	if contentType == appJson {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	token, err := h.app.LoginUsers(User)

	if err != nil {
		helpers.CustomBadRequest(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
