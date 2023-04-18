package services

import "go-rest-api-mock/models"

type ProductsServices interface {
	CreateProduct(req models.Product) (res models.Product, err error)
	GetAllProducts(userId uint, isAdmin bool) (res []models.Product, err error)
	GetProductById(productId uint) (res models.Product, err error)
	UpdateProductById(productId uint, req models.Product) (res models.Product, err error)
	DeleteProductById(id uint) (err error)
}

func (s *Services) CreateProduct(req models.Product) (res models.Product, err error) {
	res, err = s.repo.AddProduct(req)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Services) GetAllProducts(userId uint, isAdmin bool) (res []models.Product, err error) {
	if !isAdmin {
		res, err = s.repo.GetAllProducts(userId)
		if err != nil {
			return res, err
		}
		return res, nil
	}
	res, err = s.repo.GetAllProductsAdmin()
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Services) GetProductById(productId uint) (res models.Product, err error) {
	res, err = s.repo.GetProductById(productId)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Services) UpdateProductById(productId uint, req models.Product) (res models.Product, err error) {
	res, err = s.repo.UpdateProductById(productId, req)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Services) DeleteProductById(id uint) (err error) {
	err = s.repo.DeleteProductById(id)
	return err
}
