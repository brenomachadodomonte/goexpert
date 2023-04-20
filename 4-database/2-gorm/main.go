package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dns := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	//db.Create(&Product{
	//	Name:  "Notebook",
	//	Price: 1099.99,
	//})

	//products := []Product{
	//	{Name: "Notebook Samsung", Price: 1000.00},
	//	{Name: "Macbook", Price: 5400.00},
	//	{Name: "Smartphone", Price: 899.00},
	//}
	//db.Create(products)

	//var product Product
	//db.First(&product, 1)
	//fmt.Println(product)

	//var products []Product
	//db.Limit(2).Offset(2).Find(&products)
	//for _, product := range products {
	//	fmt.Println(product)
	//}

	var products []Product
	db.Where("name like ?", "%book%").Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}
}
