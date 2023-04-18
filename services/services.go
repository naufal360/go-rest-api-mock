package services

import "go-rest-api-mock/repository"

type Services struct {
	repo repository.RepoInterface
}

type ServicesInterface interface {
	UsersServices
	ProductsServices
}

func NewService(repo repository.RepoInterface) ServicesInterface {
	return &Services{repo: repo}
}
