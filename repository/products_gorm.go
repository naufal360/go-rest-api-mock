package repository

import "go-rest-api-mock/models"

type ProductsGorm interface {
	AddProduct(req models.Product) (res models.Product, err error)
	GetAllProducts(userId uint) (res []models.Product, err error)
	GetAllProductsAdmin() (res []models.Product, err error)
	GetProductById(productId uint) (res models.Product, err error)
	GetUserIdProductById(userId uint) (res models.Product, err error)
	UpdateProductById(productId uint, req models.Product) (res models.Product, err error)
	DeleteProductById(id uint) (err error)
}

func (r Repo) AddProduct(req models.Product) (res models.Product, err error) {
	err = r.gorm.Debug().Create(&req).Scan(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r Repo) GetAllProducts(userId uint) (res []models.Product, err error) {
	err = r.gorm.Model(&res).Where("user_id = ?", userId).Find(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r Repo) GetAllProductsAdmin() (res []models.Product, err error) {
	err = r.gorm.Model(&res).Find(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r Repo) GetProductById(productId uint) (res models.Product, err error) {
	err = r.gorm.Model(&res).Where("id = ?", productId).Find(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r Repo) GetUserIdProductById(userId uint) (res models.Product, err error) {
	err = r.gorm.Select("user_id").First(&res, userId).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r Repo) UpdateProductById(productId uint, req models.Product) (res models.Product, err error) {
	err = r.gorm.Model(&req).Where("id = ?", productId).Updates(
		models.Product{
			Title:       req.Title,
			Description: req.Description,
		},
	).Scan(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r Repo) DeleteProductById(id uint) (err error) {
	err = r.gorm.Model(&models.Product{}).Where("id = ?", id).Delete(&models.Product{}).Error
	if err != nil {
		return err
	}
	return nil
}
