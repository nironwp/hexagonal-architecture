package cli_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/nironwp/hexagonal-architecture/adapter/cli"
	"github.com/nironwp/hexagonal-architecture/application"
	mock_application "github.com/nironwp/hexagonal-architecture/application/mocks"
	"github.com/stretchr/testify/assert"
)

func TestRun_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productId := uuid.NewString()
	productName := "Candy"
	productPrice := 1.0
	productQuantity := 1000
	productStatus := application.ENABLED
	productMock := mock_application.NewMockProductInterface(ctrl)

	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetQuantity().Return(productQuantity).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice, productQuantity).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpect := fmt.Sprintf(
		"Product ID %s with the name %s has been created with price %f, quantity of %b and status %s",
		productMock.GetID(), productMock.GetName(), productMock.GetPrice(), productMock.GetQuantity(), productMock.GetStatus(),
	)

	result, err := cli.Run(service, "create", "", productName, productPrice, productQuantity)

	assert.Nil(t, err)
	assert.Equal(t, resultExpect, result)
}

func TestRun_Enable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productId := uuid.NewString()
	productName := "Candy"
	productPrice := 1.0
	productQuantity := 1000
	productStatus := application.ENABLED
	productMock := mock_application.NewMockProductInterface(ctrl)

	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetQuantity().Return(productQuantity).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice, productQuantity).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpect := fmt.Sprintf("Product with ID %s, and name %s is now %s", productMock.GetID(), productMock.GetName(), application.ENABLED)

	result, err := cli.Run(service, "enable", productId, "", 1.0, 1)

	assert.Nil(t, err)
	assert.Equal(t, resultExpect, result)
}

func TestRun_Disable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productId := uuid.NewString()
	productName := "Candy"
	productPrice := 1.0
	productQuantity := 1000
	productStatus := application.DISABLED
	productMock := mock_application.NewMockProductInterface(ctrl)

	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetQuantity().Return(productQuantity).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice, productQuantity).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()
	resultExpect := fmt.Sprintf("Product with ID %s, and name %s is now %s", productMock.GetID(), productMock.GetName(), application.DISABLED)

	result, err := cli.Run(service, "disable", productId, "", 1.0, 1)

	assert.Nil(t, err)
	assert.Equal(t, resultExpect, result)
}


func TestRun_Default(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productId := uuid.NewString()
	productName := "Candy"
	productPrice := 1.0
	productQuantity := 1000
	productStatus := application.ENABLED
	productMock := mock_application.NewMockProductInterface(ctrl)

	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetQuantity().Return(productQuantity).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice, productQuantity).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpect := fmt.Sprintf("Product ID: %s\n Name: %s\n Price: %f\n Quantity: %b\n Status: %s\n ", productMock.GetID(), productMock.GetName(), productMock.GetPrice(), productMock.GetQuantity(), productMock.GetStatus())

	result, err := cli.Run(service, "", productId, "", 1.0, 1)

	assert.Nil(t, err)
	assert.Equal(t, resultExpect, result)
}
