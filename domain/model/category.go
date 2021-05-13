package model

import (
	"github.com/ssuxue/mellow-server/internal/initdata"
)

type Category struct {
	ID   int    `gorm:"AUTO_INCREMENT" json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

func (category Category) TableName() string {
	return "pms_product_category"
}

func (category Category) Insert() int {
	create := initdata.DB.Create(&category)
	if create.Error != nil {
		return 0
	}
	return 1
}

func (category Category) FindAll() []Category {
	var categories []Category
	initdata.DB.Find(&categories)
	return categories
}

func (category Category) FindById() Category {
	initdata.DB.First(&category, category.ID)
	return category
}

func (category Category) DeleteOne() {
	initdata.DB.Delete(category)
}
