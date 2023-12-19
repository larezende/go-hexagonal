package cli_test

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go-hexagonal/adapters/cli"
	"go-hexagonal/application"
	mock_application "go-hexagonal/application/mocks"
	"testing"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	name := "Product Test"
	price := 10.0
	status := "enabled"
	id := "abc"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetName().Return(name).AnyTimes()
	productMock.EXPECT().GetPrice().Return(price).AnyTimes()
	productMock.EXPECT().GetStatus().Return(status).AnyTimes()
	productMock.EXPECT().GetID().Return(id).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().List().Return([]application.ProductInterface{productMock, productMock}, nil).AnyTimes()
	service.EXPECT().Create(name, price).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(id).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	expectedResult := fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s", id, name, price, status)
	result, err := cli.Run(service, "create", id, name, price)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf("Product %s %s has been enabled.", id, name)
	result, err = cli.Run(service, "enable", id, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf("Product %s %s has been disabled.", id, name)
	result, err = cli.Run(service, "disable", id, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s", id, name, price, status)
	result, err = cli.Run(service, "get", id, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s\n\nProduct ID: %s\nName: %s\nPrice: %f\nStatus: %s\n\n", id, name, price, status, id, name, price, status)
	result, err = cli.Run(service, "list", id, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)
}
