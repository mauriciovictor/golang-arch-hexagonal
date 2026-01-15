package cli_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mauriciovictor/curso-hexagonal/adapters/cli"
	mock_application "github.com/mauriciovictor/curso-hexagonal/application/mocks"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	productName := "Product Test"
	productPrice := 100.0
	productStatus := "enabled"
	productId := "abc"

	//mockando entidade de produto
	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	//mockando meu ProductService
	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(productMock).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(productMock).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID %s with the name %s  with price %f and status %s  was created successfully", productId, productName, productPrice, productStatus)

	result, err := cli.Run(service, "create", productId, productName, productPrice)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	//enabled
	resultExpected = fmt.Sprintf("Product ID %s with the name %s and price %f was enabled successfully", productId, productName, productPrice)
	result, err = cli.Run(service, "enable", productId, "", 0.0)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	//disabled
	resultExpected = fmt.Sprintf("Product ID %s with the name %s and price %f was disabled successfully", productId, productName, productPrice)
	result, err = cli.Run(service, "disable", productId, "", 0.0)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	//get
	resultExpected = fmt.Sprintf("Product ID %s with the name %s and price %f was found successfully", productId, productName, productPrice)
	result, err = cli.Run(service, "", productId, "Product Test", 0.0)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}
