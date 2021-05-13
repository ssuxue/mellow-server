package model

import (
	"github.com/ssuxue/mellow-server/internal/initdata"
)

type Product struct {
	ID                  int    `gorm:"AUTO_INCREMENT" json:"id"`
	ProductCategoryId   int64  `json:"product_category_id"`
	ProductAttributeIds string `json:"product_attribute_ids"`
	Name                string `json:"name"`
	Picture             string `json:"picture"`
	NewStatus           int8   `json:"new_status"`
	RecommendStatus     int8   `json:"recommend_status"`
	VerifyStatus        int8   `json:"verify_status"`
	Sale                int    `json:"sale"`
	Price               int    `json:"price"`
	Description         string `json:"description"`
	AlbumPictures       string `gorm:"column:album_pics" json:"album_pics"`
}

func (product Product) TableName() string {
	return "pms_product"
}

func (product Product) Insert() int {
	create := initdata.DB.Create(&product)
	if create.Error != nil {
		return 0
	}
	return 1
}

func (product Product) FindAll() []Product {
	var desserts []Product
	initdata.DB.Find(&desserts)
	return desserts
}

func (product Product) FindById() Product {
	initdata.DB.First(&product, product.ID)
	return product
}

func (product Product) FindByCategoryId(id int) []Product {
	var milkyTeas []Product
	initdata.DB.Where("product_category_id = ?", id).Find(&milkyTeas)
	return milkyTeas
}

func (product Product) DeleteOne() {
	initdata.DB.Delete(product)
}
