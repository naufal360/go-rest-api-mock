package services

import (
	"errors"
	"go-rest-api-mock/helpers"
	"go-rest-api-mock/models"
)

type UsersServices interface {
	RegisterUsers(req models.User) (res models.User, err error)
	LoginUsers(req models.User) (token string, err error)
}

func (s *Services) RegisterUsers(req models.User) (res models.User, err error) {
	res, err = s.repo.AddUsers(req)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Services) LoginUsers(req models.User) (token string, err error) {
	password := req.Password
	data, err := s.repo.AuthUsers(req)
	if err != nil {
		return "", err
	}

	comparePassword := helpers.ComparePass(data.Password, password)
	if !comparePassword {
		return "", errors.New("invalid email/password")
	}

	token = helpers.GenerateToken(req.ID, req.Email, req.Admin)

	return token, nil
}
