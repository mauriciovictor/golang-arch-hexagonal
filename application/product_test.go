package application_test

import (
	"testing"

	"github.com/mauriciovictor/curso-hexagonal/application"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestApplicationProduct_Enable(t *testing.T) {
	product := &application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()

	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "The Price must greater than zero to enable the product", err.Error())
}

func TestApplicationProduct_Disable(t *testing.T) {
	product := &application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 0

	err := product.Disable()

	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "The Price must be zero to disable the product", err.Error())
}

func TestApplicationProduct_IsValid(t *testing.T) {
	product := &application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	//product is valid
	_, err := product.IsValid()
	require.Nil(t, err)

	//status invaliid
	product.Status = "invalid"
	_, err = product.IsValid()
	require.Equal(t, "Invalid Status", err.Error())

	//price invalid
	product.Status = application.ENABLED
	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "Invalid Price", err.Error())
}
