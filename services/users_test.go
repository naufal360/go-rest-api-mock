package services

import (
	"go-rest-api-mock/models"
	"go-rest-api-mock/services/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterUsers(t *testing.T) {
	mockData := new(mocks.GormMock)

	data_one := models.User{
		Fullname: "ahmad",
		Email:    "ahmad@gmail.com",
		Password: "password123",
	}

	tests := []struct {
		name         string
		userId       int
		code         int
		body         models.User
		expectedBody models.User
	}{
		{
			name:   "success found",
			userId: 1,
			body: models.User{
				Fullname: "ahmad",
				Email:    "ahmad@gmail.com",
				Password: "password123",
			},
			expectedBody: data_one,
		},
		{
			name:         "bad request",
			userId:       2,
			body:         models.User{},
			expectedBody: models.User{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockData.Mock.On("AddUsers", test.body).Return(test.body, nil)
			services := NewService(mockData)

			data, err := services.RegisterUsers(test.body)

			assert.Nil(t, err)
			if err != nil {
				assert.NotEqual(t, test.expectedBody, data)
			} else {
				assert.Equal(t, test.expectedBody, data)
			}
		})
	}
}
