package mocks

import "github.com/stretchr/testify/mock"

type GormMock struct {
	Mock mock.Mock
}

type AllMocks interface {
	UsersRepoInter
	ProductsRepoInter
}
