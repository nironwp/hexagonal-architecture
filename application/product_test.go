package application

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductCreation_When_all_Valid_Parameters(t *testing.T) {

	dto := NewProductDto{
		Name:     "Candy",
		Price:    1,
		Quantity: 1000,
	}

	_, err := NewProduct(dto)

	assert.Nil(t, err)
}

func TestShouldReturnAnErrorWhenNewProduct_WithInvalidName(t *testing.T) {
	dto := NewProductDto{
		Name:     "",
		Price:    1,
		Quantity: 1000,
	}

	_, err := NewProduct(dto)

	assert.Equal(t, "Invalid name", err.Error())
}

func TestShouldReturnAnErrorWhenNewProduct_WithInvalidPrice(t *testing.T) {
	dto := NewProductDto{
		Name:     "Candy",
		Price:    -1,
		Quantity: 1000,
	}

	_, err := NewProduct(dto)

	assert.Equal(t, "Invalid price", err.Error())
}

func TestShouldReturnAnErrorWhenNewProduct_WithInvalidQuantity(t *testing.T) {
	dto := NewProductDto{
		Name:     "Candy",
		Price:    1,
		Quantity: -1,
	}

	_, err := NewProduct(dto)

	assert.Equal(t, "Invalid quantity", err.Error())
}

func TestShoudEnableProduct(t *testing.T) {
	assert := assert.New(t)
	dto := NewProductDto{
		Name:     "Candy",
		Price:    1,
		Quantity: 1000,
	}

	product, err := NewProduct(dto)

	assert.Nil(err)

	err = product.Enable()

	assert.Nil(err)
	assert.Equal(ENABLED, product.Status)
}

func TestShoudDesableProduct(t *testing.T) {
	assert := assert.New(t)
	dto := NewProductDto{
		Name:     "Candy",
		Price:    1,
		Quantity: 1000,
	}

	product, err := NewProduct(dto)

	assert.Nil(err)

	err = product.Enable()

	assert.Nil(err)
	assert.Equal(ENABLED, product.Status)

	product.Disable()

	assert.Equal(DISABLED, product.Status)
}

func TestShouldReturnAStatusOfProduct(t *testing.T) {
	assert := assert.New(t)
	dto := NewProductDto{
		Name:     "Candy",
		Price:    1,
		Quantity: 1000,
	}

	product, err := NewProduct(dto)

	assert.Nil(err)

	assert.Equal(product.Status, DISABLED)
}

func TestShouldReturnAStatusOfProductWhenIsEnabled(t *testing.T) {
	assert := assert.New(t)
	dto := NewProductDto{
		Name:     "Candy",
		Price:    1,
		Quantity: 1000,
	}

	product, err := NewProduct(dto)

	assert.Nil(err)

	err = product.Enable()

	assert.Nil(err)

	assert.Equal(product.Status, ENABLED)
}

func TestSholdReturnAnName(t *testing.T) {
	assert := assert.New(t)
	dto := NewProductDto{
		Name:     "Candy",
		Price:    1,
		Quantity: 1000,
	}

	product, err := NewProduct(dto)
	assert.Nil(err)

	assert.Equal(product.GetName(), "Candy")
}

func TestSholdReturnAnPrice(t *testing.T) {
	assert := assert.New(t)
	dto := NewProductDto{
		Name:     "Candy",
		Price:    1,
		Quantity: 1000,
	}

	product, err := NewProduct(dto)
	assert.Nil(err)

	assert.Equal(product.GetPrice(), 1.0)
}

func TestSholdReturnAnQuantity(t *testing.T) {
	assert := assert.New(t)
	dto := NewProductDto{
		Name:     "Candy",
		Price:    1,
		Quantity: 1000,
	}

	product, err := NewProduct(dto)
	assert.Nil(err)

	assert.Equal(product.GetQuantity(), 1000)
}

func TestShouldChangePrice(t *testing.T) {
	assert := assert.New(t)
	dto := NewProductDto{
		Name:     "Candy",
		Price:    1,
		Quantity: 1000,
	}

	product, err := NewProduct(dto)
	assert.Nil(err)

	err = product.ChangePrice(2.0)

	assert.Nil(err, nil)
	assert.Equal(2.0, product.GetPrice())
}

func TestShouldTrowErrorWhenChangeInvalidPrice(t *testing.T) {
	assert := assert.New(t)
	dto := NewProductDto{
		Name:     "Candy",
		Price:    1,
		Quantity: 1000,
	}

	product, err := NewProduct(dto)
	assert.Nil(err)

	err = product.ChangePrice(-1.0)

	assert.Equal("Invalid price", err.Error())
}

func TestShouldChangeQuantity(t *testing.T) {
	assert := assert.New(t)
	dto := NewProductDto{
		Name:     "Candy",
		Price:    1,
		Quantity: 1000,
	}

	product, err := NewProduct(dto)
	assert.Nil(err)

	err = product.ChangeQuantity(500)
	assert.Nil(err)
	assert.Equal(500, product.GetQuantity())
}

func TestShouldTrowErrorWhenChangeInvalidQuantity(t *testing.T) {
	assert := assert.New(t)
	dto := NewProductDto{
		Name:     "Candy",
		Price:    1,
		Quantity: 1000,
	}

	product, err := NewProduct(dto)
	assert.Nil(err)

	err = product.ChangeQuantity(-100)

	assert.Equal("Invalid quantity", err.Error())
}
