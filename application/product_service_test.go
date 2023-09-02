package application_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go-hexagonal/application"
	applicationMocks "go-hexagonal/application/mocks"
	"testing"
)

func TestProductService_Get(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productMock := applicationMocks.NewMockProductInterface(ctrl)
	persistenceMock := applicationMocks.NewMockProductPersistenceInterface(ctrl)
	persistenceMock.EXPECT().Get(gomock.Any()).Return(productMock, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistenceMock,
	}

	result, err := service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, result, productMock)
}

func TestProductService_Save(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productMock := applicationMocks.NewMockProductInterface(ctrl)
	persistenceMock := applicationMocks.NewMockProductPersistenceInterface(ctrl)
	persistenceMock.EXPECT().Save(gomock.Any()).Return(productMock, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistenceMock,
	}

	result, err := service.Create("Product 1", 10)
	require.Nil(t, err)
	require.Equal(t, result, productMock)
}

func TestProductService_EnableDisable(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productMock := applicationMocks.NewMockProductInterface(ctrl)
	productMock.EXPECT().Enable().Return(nil).AnyTimes()
	productMock.EXPECT().Disable().Return(nil).AnyTimes()
	persistenceMock := applicationMocks.NewMockProductPersistenceInterface(ctrl)
	persistenceMock.EXPECT().Save(gomock.Any()).Return(productMock, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistenceMock,
	}

	result, err := service.Enable(productMock)
	require.Nil(t, err)
	require.Equal(t, result, productMock)

	result, err = service.Disable(productMock)
	require.Nil(t, err)
	require.Equal(t, result, productMock)

}
