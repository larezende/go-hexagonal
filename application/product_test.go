package application_test

import (
	uuid "github.com/satori/go.uuid"
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

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()

	require.Nil(t, err)
	require.Equal(t, application.DISABLED, product.Status)

	product.Price = 10

	err = product.Disable()

	require.NotNil(t, err)
	require.Equal(t, "the price must be zero in order to have the product disabled", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()

	require.Nil(t, err)

	product.Name = ""
	_, err = product.IsValid()

	require.Equal(t, "Name: non zero value required", err.Error())

	product.Name = "Hello"
	product.Status = "invalid"
	_, err = product.IsValid()

	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.DISABLED
	product.Price = -10
	_, err = product.IsValid()

	require.Equal(t, "the price must be greater or equal 0", err.Error())

	product.Price = 10
	_, err = product.IsValid()

	product.ID = "123"
	_, err = product.IsValid()

	require.Equal(t, "ID: 123 does not validate as uuidv4", err.Error())

	product.ID = uuid.NewV4().String()
	_, err = product.IsValid()
	require.Nil(t, err)
}
