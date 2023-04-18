package repository

import "go-rest-api-mock/models"

type UsersGorm interface {
	AddUsers(req models.User) (res models.User, err error)
	AuthUsers(req models.User) (res models.User, err error)
}

func (r Repo) AddUsers(req models.User) (res models.User, err error) {
	err = r.gorm.Debug().Create(&req).Scan(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r Repo) AuthUsers(req models.User) (res models.User, err error) {
	err = r.gorm.Debug().Where("email = ?", req.Email).Take(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}
