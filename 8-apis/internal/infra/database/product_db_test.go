package database

import (
	"fmt"
	"github.com/brenomachadodomonte/goexpert/apis/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

func TestProduct_Create(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Macbook", 5000.0)
	assert.Nil(t, err)

	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.Nil(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestProduct_FindAll(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64())
		assert.Nil(t, err)
		db.Create(product)
	}

	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[2].Name)
}

func TestProduct_FindByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Macbook", 5000.0)
	assert.Nil(t, err)

	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.Nil(t, err)
	assert.NotEmpty(t, product.ID)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, productFound.ID, product.ID)
	assert.Equal(t, productFound.Name, product.Name)
	assert.Equal(t, productFound.Price, product.Price)
}

func TestProduct_Update(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Macbook", 5000.0)
	assert.Nil(t, err)

	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.Nil(t, err)
	assert.NotEmpty(t, product.ID)

	product.Name = "Notebook"
	product.Price = 4500.0

	err = productDB.Update(product)
	assert.Nil(t, err)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, productFound.ID, product.ID)
	assert.Equal(t, productFound.Name, "Notebook")
	assert.Equal(t, productFound.Price, 4500.0)
}

func TestProduct_Delete(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Macbook", 5000.0)
	assert.Nil(t, err)

	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.Nil(t, err)
	assert.NotEmpty(t, product.ID)

	err = productDB.Delete(product.ID.String())
	assert.Nil(t, err)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.Nil(t, productFound)
	assert.NotNil(t, err)

}
