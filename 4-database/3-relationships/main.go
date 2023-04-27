package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dns := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	////create category
	//category := &Category{Name: "Kitchen"}
	//db.Create(category)
	//
	////create Product
	//product := &Product{
	//	Name:       "Fry pan",
	//	Price:      7.10,
	//	CategoryID: category.ID,
	//}
	//db.Create(product)
	//
	//db.Create(&SerialNumber{
	//	Number:    "123456",
	//	ProductID: product.ID,
	//})

	//var products []Product
	//db.Preload("Category").Preload("SerialNumber").Find(&products)
	//for _, product := range products {
	//	fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	//}

	var categories []Category
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("- ", product.Name, product.SerialNumber.Number)
		}
	}
}
