package routes

import (
	"go-rest-api-mock/handlers"
	"go-rest-api-mock/middlewares"
	"go-rest-api-mock/services"

	"github.com/gin-gonic/gin"
)

func RegisterAPI(r *gin.Engine, app services.ServicesInterface) {
	server := handlers.NewHttpServer(app)

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", server.UserRegister)
		userRouter.POST("/login", server.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", server.CreateProduct)
		productRouter.GET("/", server.GetAllProduct)
		productRouter.GET("/:productId", middlewares.ProductAuthorization(), server.GetProductById)
		productRouter.PUT("/:productId", middlewares.ProductUpdateDeleteAuthorization(), server.UpdateProduct)
		productRouter.DELETE("/:productId", middlewares.ProductUpdateDeleteAuthorization(), server.DeleteProductById)
	}
}
