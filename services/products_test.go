package services

import (
	"go-rest-api-mock/models"
	"go-rest-api-mock/services/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	mockData := new(mocks.GormMock)

	data_one := models.Product{
		Title:       "book1",
		Description: "haha1",
	}

	tests := []struct {
		name         string
		userId       int
		code         int
		body         models.Product
		expectedBody models.Product
	}{
		{
			name:   "success found",
			userId: 1,
			body: models.Product{
				Title:       "book1",
				Description: "haha1",
			},
			expectedBody: data_one,
		},
		{
			name:         "bad request",
			userId:       2,
			body:         models.Product{},
			expectedBody: models.Product{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockData.Mock.On("AddProduct", test.body).Return(test.body, nil)
			services := NewService(mockData)

			data, err := services.CreateProduct(test.body)

			assert.Nil(t, err)
			if err != nil {
				assert.NotEqual(t, test.expectedBody, data)
			} else {
				assert.Equal(t, test.expectedBody, data)
			}
		})
	}
}

func TestGetAllProducts(t *testing.T) {
	mockData := new(mocks.GormMock)

	data_one := models.Product{
		Title:       "book1",
		Description: "haha1",
		UserID:      1,
	}
	data_two := models.Product{
		Title:       "book2",
		Description: "haha2",
		UserID:      1,
	}

	tests := []struct {
		name         string
		userId       int
		code         int
		body         []models.Product
		expectedBody []models.Product
	}{
		{
			name:   "success found",
			userId: 1,
			body: []models.Product{
				{
					Title:       "book1",
					Description: "haha1",
					UserID:      1,
				},
				{
					Title:       "book2",
					Description: "haha2",
					UserID:      1,
				},
			},
			expectedBody: []models.Product{
				data_one,
				data_two,
			},
		},
		{
			name:         "failed not found",
			userId:       2,
			body:         nil,
			expectedBody: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockData.Mock.On("GetAllProducts", uint(test.userId)).Return(test.body, nil)
			services := NewService(mockData)

			isAdmin := false
			data, err := services.GetAllProducts(uint(test.userId), isAdmin)

			assert.Nil(t, err)
			assert.Equal(t, test.expectedBody, data)
		})
	}
}
func TestGetProductById(t *testing.T) {
	mockData := new(mocks.GormMock)

	data := models.Product{
		GormModel: models.GormModel{
			ID: 1,
		},
		Title:       "book1",
		Description: "haha1",
	}

	tests := []struct {
		name         string
		id           int
		code         int
		body         models.Product
		expectedBody models.Product
	}{
		{
			name:         "success found",
			id:           1,
			body:         data,
			expectedBody: data,
		},
		{
			name:         "failed not found",
			id:           3,
			body:         models.Product{},
			expectedBody: models.Product{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockData.Mock.On("GetProductById", uint(test.id)).Return(test.body, nil)
			services := NewService(mockData)

			data, err := services.GetProductById(uint(test.id))

			assert.Nil(t, err)
			assert.Equal(t, test.expectedBody, data)
		})
	}
}

func TestUpdateProductById(t *testing.T) {
	mockData := new(mocks.GormMock)

	data := models.Product{
		GormModel: models.GormModel{
			ID: 1,
		},
		Title:       "book1",
		Description: "haha1",
	}

	tests := []struct {
		name         string
		id           int
		code         int
		body         models.Product
		expectedBody models.Product
	}{
		{
			name:         "success found",
			id:           1,
			body:         data,
			expectedBody: data,
		},
		{
			name:         "failed not found",
			id:           3,
			body:         models.Product{},
			expectedBody: models.Product{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockData.Mock.On("UpdateProductById", uint(test.id), test.body).Return(test.body, nil)
			services := NewService(mockData)

			data, err := services.UpdateProductById(uint(test.id), test.body)

			assert.Nil(t, err)
			assert.Equal(t, test.expectedBody, data)
		})
	}
}

func TestDeleteProductById(t *testing.T) {
	mockData := new(mocks.GormMock)

	tests := []struct {
		name         string
		id           int
		code         int
		body         map[string]string
		expectedBody map[string]string
	}{
		{
			name: "success found",
			id:   1,
			body: map[string]string{
				"message": "Success Delete Product",
			},
			expectedBody: map[string]string{
				"message": "Success Delete Product",
			},
		},
		{
			name: "failed not found",
			id:   3,
			body: map[string]string{
				"error":   "Data Not Found",
				"message": "data doesn't exist",
			},
			expectedBody: map[string]string{
				"error":   "Data Not Found",
				"message": "data doesn't exist",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockData.Mock.On("DeleteProductById", uint(test.id)).Return(nil)
			services := NewService(mockData)

			err := services.DeleteProductById(uint(test.id))

			assert.Nil(t, err)
			if err != nil {
				assert.NotEqual(t, test.expectedBody, test.body)
			} else {
				assert.Equal(t, test.expectedBody, test.body)
			}
		})
	}
}
