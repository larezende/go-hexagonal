package application_test

import (
	"github.com/stretchr/testify/require"
	"go-hexagonal/application"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()

	require.Nil(t, err)
	require.Equal(t, application.ENABLED, product.Status)

	product.Price = 0

	err = product.Enable()

	require.NotNil(t, err)
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}
