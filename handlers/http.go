package handlers

import "go-rest-api-mock/services"

type HttpServer struct {
	app services.ServicesInterface
}

func NewHttpServer(app services.ServicesInterface) HttpServer {
	return HttpServer{app: app}
}
