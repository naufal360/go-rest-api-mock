package app

import (
	"fmt"
	"go-rest-api-mock/config"
	"go-rest-api-mock/repository"
	"go-rest-api-mock/routes"
	"go-rest-api-mock/services"
	"os"

	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApp() {
	repo := repository.NewRepo(config.GORM.DB)
	app := services.NewService(repo)
	routes.RegisterAPI(router, app)

	port := os.Getenv("APP_PORT")
	router.Use(gin.Recovery())
	router.Run(fmt.Sprintf(":%s", port))
}
