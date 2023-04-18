package mocks

import (
	"go-rest-api-mock/models"
)

type ProductsRepoInter interface {
	AddProduct(req models.Product) (res models.Product, err error)
	GetAllProducts(userId uint) (res []models.Product, err error)
	GetAllProductsAdmin() (res []models.Product, err error)
	GetProductById(productId uint) (res models.Product, err error)
	GetUserIdProductById(userId uint) (res models.Product, err error)
	UpdateProductById(productId uint, req models.Product) (res models.Product, err error)
	DeleteProductById(id uint) (err error)
}

func (repo *GormMock) AddProduct(req models.Product) (res models.Product, err error) {
	arguments := repo.Mock.Called(req)

	if arguments.Get(0) == nil {
		return res, nil
	}

	res = arguments.Get(0).(models.Product)
	err = arguments.Error(1)

	return res, err
}

func (repo *GormMock) GetAllProducts(userId uint) (res []models.Product, err error) {
	arguments := repo.Mock.Called(userId)

	if arguments.Get(0) == nil {
		return res, nil
	}

	res = arguments.Get(0).([]models.Product)
	err = arguments.Error(1)

	return res, err
}

func (repo *GormMock) GetAllProductsAdmin() (res []models.Product, err error) {
	arguments := repo.Mock.Called()

	if arguments.Get(0) == nil {
		return res, nil
	}

	res = arguments.Get(0).([]models.Product)
	err = arguments.Error(1)

	return res, err
}

func (repo *GormMock) GetProductById(productId uint) (res models.Product, err error) {
	arguments := repo.Mock.Called(productId)

	if arguments.Get(0) == nil {
		return res, nil
	}

	res = arguments.Get(0).(models.Product)
	err = arguments.Error(1)

	return res, err
}

func (repo *GormMock) GetUserIdProductById(userId uint) (res models.Product, err error) {
	arguments := repo.Mock.Called(userId)

	if arguments.Get(0) == nil {
		return res, nil
	}

	res = arguments.Get(0).(models.Product)
	err = arguments.Error(1)

	return res, err
}
func (repo *GormMock) UpdateProductById(productId uint, req models.Product) (res models.Product, err error) {
	arguments := repo.Mock.Called(req)

	if arguments.Get(0) == nil {
		return res, nil
	}

	res = arguments.Get(0).(models.Product)
	err = arguments.Error(1)

	return res, err
}

func (repo *GormMock) DeleteProductById(id uint) (err error) {
	arguments := repo.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	}

	err = arguments.Error(1)

	return err
}
