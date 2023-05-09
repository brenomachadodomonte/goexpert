package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Macbook", 10.0)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, "Macbook", product.Name)
	assert.Equal(t, 10.0, product.Price)
}

func TestNewProductWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("", 10.0)
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Error(t, ErrNameIsRequired)
}

func TestNewProductWhenPriceIsRequired(t *testing.T) {
	product, err := NewProduct("Macbook", 0.0)
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Error(t, ErrPriceIsRequired)
}

func TestNewProductWhenPriceIsInvalid(t *testing.T) {
	product, err := NewProduct("Macbook", -10.0)
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Error(t, ErrInvalidPrice)
}
