package mocks

import "go-rest-api-mock/models"

type UsersRepoInter interface {
	AddUsers(req models.User) (res models.User, err error)
	AuthUsers(req models.User) (res models.User, err error)
}

func (repo *GormMock) AddUsers(req models.User) (res models.User, err error) {
	arguments := repo.Mock.Called(req)

	if arguments.Get(0) == nil {
		return res, nil
	}

	res = arguments.Get(0).(models.User)
	err = arguments.Error(1)

	return res, err
}

func (repo *GormMock) AuthUsers(req models.User) (res models.User, err error) {
	arguments := repo.Mock.Called(req)

	if arguments.Get(0) == nil {
		return res, nil
	}

	res = arguments.Get(0).(models.User)
	err = arguments.Error(1)

	return res, err
}
